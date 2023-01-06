package main

import (
	"fmt"

	"lem-in/class"
	decryptdata "lem-in/decryptData"
	"lem-in/errorsLem"
	findpath "lem-in/findPath"
)

func main() {
	farm, err := decryptdata.ParseData()
	if err != nil {
		panic(err)
	}
	// for _, room := range farm.Rooms {
	// 	fmt.Print(room.Name, ": ")
	// 	for _, link := range room.RoomsLink {
	// 		fmt.Print(link.Name, " ")
	// 	}
	// 	fmt.Println()
	// }
	res := findpath.GetBestPath(farm)
	for _, path := range *res {
		fmt.Print("Path: ")
		for _, roomPath := range path {
			fmt.Print(roomPath.Name, " -> ")
		}
		fmt.Println()
	}

	// var maxPath int
	// if l1, l2 := len(farm.End.RoomsLink), len(farm.Start.RoomsLink); l1 > l2 {
	// 	maxPath = l2
	// } else {
	// 	maxPath = l1
	// }

	test := findpath.FindNoCrossPathList(*res, *farm)
	for _, x := range test {
		for _, y := range x {
			for _, z := range y {
				fmt.Print(z.Name)
				fmt.Print(" -> ")
			}
			fmt.Println("")
		}
		fmt.Println()
	}
	bestPath := [][]*class.Room{}
	shortestEtape := 0
	for _, pathLs := range test {
		n := farm.AntNb
		table := []int{}
		for _, p := range pathLs {
			table = append(table, len(p)-1)
		}
		fmt.Println(table)
		for n != 0 {
			table[findMin(table)]++
			n--
		}
		fmt.Println(table)
		fmt.Println()
		if newLen := findMaxValue(table); shortestEtape > newLen || shortestEtape == 0 {
			shortestEtape = newLen
			bestPath = pathLs
		}
	}
	if len(bestPath) == 0 {
		panic(errorsLem.ErrNoPathFound)
	}
	fmt.Println("nb etape =", shortestEtape-1, (bestPath))

	pathObj := []*class.PathO{}
	for i, p := range bestPath {
		pathObj = append(pathObj, &class.PathO{Name: fmt.Sprint("path", i), Path: p, IsUsed: false, Population: len(p)})
	}
	fmt.Println(pathObj)
	farm.CreatePopulation()

	popRunI := 1
	for i := 0; i < shortestEtape; i++ {
		for range bestPath {
			if popRunI < farm.AntNb {
				popRunI++
			}
		}
		for _, ant := range farm.Population[:popRunI] {
			if len(ant.Path.Path) == 0 {
				iM := findMinPath(pathObj)
				for pathObj[iM].IsUsed {
					iM = findMinPath(append(pathObj[:iM], pathObj[1+iM:]...))
				}
				ant.Path = pathObj[iM]
				pathObj[iM].Population++
			}
			// fmt.Print(ant.ID)
		}
		for _, p := range pathObj {
			p.IsUsed = false
		}
		// for range bestPath {
		// 	if len(farm.Population) != farm.AntNb {
		// 		farm.AddPopulation()
		// 		iM := findMin(pathPopulation)
		// 		farm.Population[len(farm.Population)-1].Path = bestPath[iM]
		// 		pathPopulation[iM]++
		// 	}
		// }
		// fmt.Print(len(farm.Population))

		// fmt.Print(popRunI)
		for _, ant := range farm.Population[:popRunI] {
			// fmt.Print(ant.ID)
			
			ant.MoveAnt()
			if ant.Path.Path[ant.PositionI].Name == farm.End.Name {
				ant.Path.Path[ant.PositionI].DeletAnt(ant)	
			}
		}
		// fmt.Print(pathPopulation)

		// for _, txt := range tableMove {
		// 	if txt != "" {
		// 		fmt.Print(txt)
		// 	}
		// }
		fmt.Println()
	}
}

// MOD on "Alann" Branch

func findMin(table []int) int {
	minI := 0
	for i := range table {
		if table[i] < table[minI] {
			minI = i
		}
	}
	return minI
}

func findMinPath(table []*class.PathO) int {
	minI := 0
	for i := range table {
		if table[i].Population < table[minI].Population {
			minI = i
		}
	}
	return minI
}

func findMaxValue(table []int) int {
	maxV := 0
	for i := range table {
		if table[i] > maxV {
			maxV = table[i]
		}
	}
	return maxV
}

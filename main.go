package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

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
	for _, room := range farm.Rooms {
		fmt.Print(room.Name, ": ")
		for _, link := range room.RoomsLink {
			fmt.Print(link.Name, " ")
		}
		fmt.Println()
	}
	res := findpath.GetBestPath(farm)
	for i, path := range *res {
		fmt.Print("Path n°", i+1, ": ")
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
				// fmt.Print(i)
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

	stepTable := []class.Step{}

	for i := 0; i < shortestEtape; i++ {
		stepTable = append(stepTable, class.Step{Ants: []class.NewAnt{}, Paths: []class.NewRoom{}})
		for _, ant := range farm.Population[:] {
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
		for _, ant := range farm.Population[:] {
			// fmt.Print(len(ant.Path.Path) == 0 || ant.Path.Path[ant.PositionI].Name == farm.End.Name)
			if len(ant.Path.Path) == 0 || ant.Path.Path[ant.PositionI].Name == farm.End.Name {
				continue
			}
			pos := ant.PositionI
			ant.MoveAnt()
			if pos == ant.PositionI {
				continue
			}
			stepTable[len(stepTable)-1].Paths = append(stepTable[len(stepTable)-1].Paths, ant.Path.Path[ant.PositionI].GetNewRoom())
			stepTable[len(stepTable)-1].Ants = append(stepTable[len(stepTable)-1].Ants, ant.GetNewAnt())
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
	// fmt.Println(stepTable)
	jsonData, err := json.MarshalIndent(class.ToSjson(farm, stepTable), "", "	")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("ant.json", jsonData, 0o1411)
	fmt.Println(len(test))
	fmt.Println(len(*res))
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

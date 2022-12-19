package main

import (
	"fmt"

	"lem-in/class"
	decryptdata "lem-in/decryptData"
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
	fmt.Println("nb etape =", shortestEtape-1, (bestPath))
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

func findMaxValue(table []int) int {
	maxV := 0
	for i := range table {
		if table[i] > maxV {
			maxV = table[i]
		}
	}
	return maxV
}

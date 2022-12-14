package main

import (
	"fmt"

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
}

// MOD on "Alann" Branch

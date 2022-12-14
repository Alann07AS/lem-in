package findpath

import (
	"lem-in/class"
)

func GetBestPath(farm *class.Farm) *[][]*class.Room {
	allPath := &[][]*class.Room{}
	recu(farm.Start, []*class.Room{farm.Start}, farm, allPath)
	sort(*allPath)
	return allPath
}

func recu(room *class.Room, path []*class.Room, farm *class.Farm, allPath *[][]*class.Room) {
	if contains(path, farm.End) {
		pathToSave := make([]*class.Room, len(path))
		copy(pathToSave, path)
		*allPath = append(*allPath, pathToSave)
		return
	}
	for _, link := range room.RoomsLink {
		if contains(path, link) {
			continue
		}
		path = append(path, link)
		recu(link, path, farm, allPath)
		path = path[:len(path)-1]
	}
}

func contains(s []*class.Room, str *class.Room) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func sort(table [][]*class.Room) {
	isMod := true
	for isMod {
		isMod = false
		for i := 1; i < len(table); i++ {
			if len(table[i]) < len(table[i-1]) {
				isMod = true
				table[i-1], table[i] = table[i], table[i-1]
			}
		}
	}
}

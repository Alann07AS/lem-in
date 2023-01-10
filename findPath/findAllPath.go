package findpath

import (
	"lem-in/class"
)

func GetBestPath(farm *class.Farm) *[][]*class.Room {
	allPath := &[][]*class.Room{}
	recuPathFinder(farm.Start, []*class.Room{farm.Start}, farm, allPath)
	// sort(*allPath)
	return allPath
}

func recuPathFinder(room *class.Room, path []*class.Room, farm *class.Farm, allPath *[][]*class.Room) {
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
		recuPathFinder(link, path, farm, allPath)
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

func containsSlice(s [][]*class.Room, str []*class.Room) bool {
	l := len(str)
	isContain := true
	for _, path := range s {
		if len(path) != l {
			continue
		}
		for i := 0; i < l; i++ {
			if path[i] != str[i] {
				isContain = false
			}
		}
		if isContain {
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

func FindNoCrossPathList(allPath [][]*class.Room, farm class.Farm) [][][]*class.Room {
	shortestList := [][][]*class.Room{}
	for _, pathAll := range allPath {
		shortestPath := [][]*class.Room{pathAll}
		for _, a := range allPath {
			isCross := func() bool {
				for _, v := range shortestPath {
					if !isNotCross(a, v, farm) {
						return true
					}
				}
				return false
			}()
			if containsSlice(shortestPath, a) || isCross {
				continue
			}
			temp := make([]*class.Room, len(a))
			copy(temp, a)
			shortestPath = append(shortestPath, temp)
		}
		temp := make([][]*class.Room, len(shortestPath))
		copy(temp, shortestPath)
		shortestList = append(shortestList, shortestPath)
	}

	return shortestList
}

func isNotCross(p1, p2 []*class.Room, farm class.Farm) bool {
	for _, r1 := range p1 {
		if contains(p2, r1) && r1 != farm.Start && r1 != farm.End {
			return false
		}
	}
	return true
}

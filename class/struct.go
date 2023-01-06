package class

import (
	"fmt"
)

type Room struct {
	Name       string
	X          string
	Y          string
	Population []*Ant
	RoomsLink  []*Room
}

type Ant struct {
	ID        int
	Path      *PathO
	PositionI int
}

type Farm struct {
	AntNb      int
	Start      *Room
	End        *Room
	Rooms      []*Room
	Population []*Ant
}

func CreateRoom(name string, x, y string) *Room {
	return &Room{
		Name:      name,
		X:         x,
		Y:         y,
		RoomsLink: []*Room{},
	}
}

func AddLinkRoom(r1, r2 *Room) {
	r1.RoomsLink = append(r1.RoomsLink, r2)
	r2.RoomsLink = append(r2.RoomsLink, r1)
}

func (famr *Farm) CreatePopulation() {
	table := make([]*Ant, famr.AntNb)
	for i := 0; i < famr.AntNb; i++ {
		table[i] = &Ant{ID: i + 1, PositionI: 0, Path: &PathO{}}
	}
	famr.Population = table
}

func (famr *Farm) AddPopulation() {
	famr.Population = append(famr.Population, &Ant{ID: len(famr.Population) + 1, PositionI: 0})
}

func (ro *Room) IsFull() bool {
	return len(ro.Population) != 0
}

func (famr *Farm) PushAnt(path []*Room, tableMove []string) []string {
	previousAnt := 0
	for i, room := range path[:len(path)-1] {
		if len(room.Population) != 0 && previousAnt != room.Population[0].ID {
			path[i+1].Population = append(path[i+1].Population, room.Population[0])
			previousAnt = room.Population[0].ID
			tableMove[previousAnt-1] = fmt.Sprint("L", room.Population[0].ID, "-", path[i+1].Name, " ")
			room.Population = room.Population[1:]
		}
	}
	return tableMove
}

func NbAntInPath(path []*Room) int {
	nb := 0
	for _, room := range path[1 : len(path)-1] {
		nb += len(room.Population)
	}
	return nb
}

// func (a *Ant) MoveAnt() {
// 	// dest.Population = append(dest.Population, a)
// 	if !a.Path[a.PositionI+1].IsFull() || a.PositionI+1 == len(a.Path)-1 {
// 		// fmt.Print(a.ID)
// 		if len(a.Path[a.PositionI].Population) != 0 {
// 			a.Path[a.PositionI].Population = a.Path[a.PositionI].Population[1:]
// 		}
// 		a.PositionI++
// 		a.Path[a.PositionI].Population = append(a.Path[a.PositionI].Population, a)
// 		fmt.Print("L", a.ID, "-", a.Path[a.PositionI].Name, " ")
// 	}
// }

func (f *Room) DeletAnt(a *Ant) {
	for i, v := range f.Population {
		if v.ID == a.ID {
			f.Population = append(f.Population[:i], f.Population[i+1:]...)
		}
	}
}

type PathO struct {
	Name       string
	Path       []*Room
	IsUsed     bool
	Population int
}

func (a *Ant) MoveAnt() {
	// fmt.Print(a.ID, a.Path.IsUsed)
	if a.PositionI == len(a.Path.Path)-1 || (a.Path.IsUsed) {
		// La fourmi a atteint sa destination finale, elle ne bouge plus
		return
	}
	if !a.Path.Path[a.PositionI+1].IsFull() {
		// La salle suivante n'est pas pleine, on peut déplacer la fourmi
		if len(a.Path.Path[a.PositionI].Population) != 0 {
			a.Path.Path[a.PositionI].Population = a.Path.Path[a.PositionI].Population[1:]
		}
		a.PositionI++
		a.Path.Path[a.PositionI].Population = append(a.Path.Path[a.PositionI].Population, a)
		fmt.Print("L", a.ID, "-", a.Path.Path[a.PositionI].Name, " ")
	}
	if len(a.Path.Path) == 2 {
		a.Path.IsUsed = true
	}
	// fmt.Println(a.Path[a.PositionI].Name)
	// Si la salle suivante est pleine, la fourmi ne bouge pas et attend l'étape suivante
}

func (farm *Farm) GetRoomByName(name string) *Room {
	for _, room := range farm.Rooms {
		if room.Name == name {
			return room
		}
	}
	return nil
}

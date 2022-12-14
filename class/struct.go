package class

type Room struct {
	Name       string
	X          string
	Y          string
	Population []Ant
	RoomsLink  []*Room
}

type Ant struct {
	Id int
}

type Farm struct {
	AntNb int
	Start *Room
	End   *Room
	Rooms []*Room
}

func CreatRoom(name string, x, y string) *Room {
	return &Room{
		Name:       name,
		X:          x,
		Y:          y,
		Population: []Ant{},
		RoomsLink:  []*Room{},
	}
}

func AddLinkRoom(r1, r2 *Room) {
	r1.RoomsLink = append(r1.RoomsLink, r2)
	r2.RoomsLink = append(r2.RoomsLink, r1)
}

func (room *Room) IsEmpty() bool {
	if len(room.Population) != 0 {
		return false
	} else {
		return true
	}
}

func CreatPopulation(n int) []Ant {
	table := make([]Ant, n)
	for i := 0; i < n; i++ {
		table[i] = Ant{Id: i + 1}
	}
	return table
}

func (farm *Farm) GetRoomByName(name string) *Room {
	for _, room := range farm.Rooms {
		if room.Name == name {
			return room
		}
	}
	return nil
}

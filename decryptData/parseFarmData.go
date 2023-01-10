package decryptdata

import (
	"os"
	"strconv"
	"strings"

	"lem-in/class"
	"lem-in/errorsLem"
)

func ParseData() (*class.Farm, error) {
	// regNbAnt := regexp.MustCompile(`(?m)^\d\n`)
	// regStartRoom := regexp.Compile("")
	// regEndRoom := regexp.Compile("")
	// regRoom := regexp.Compile("")
	// regLink := regexp.Compile("")

	farm := &class.Farm{}
	args := os.Args
	// ckeck for name fille
	if len(args) != 2 {
		return farm, errorsLem.ErrNoArgs
	}
	data, err := os.ReadFile(args[1])
	// check for file read sucesfull
	if err != nil {
		return farm, err
	}

	// start parse data
	lines := strings.Split(string(data), "\n")
	farm.AntNb, err = strconv.Atoi(lines[0])
	if err != nil || farm.AntNb <= 0 {
		return farm, errorsLem.ErrNbLemin
	}
	for i := 0; i < len(lines); i++ {
		if lines[i] == "##start" {
			roomInfo := strings.Split(lines[i+1], " ")
			farm.Start = class.CreateRoom(roomInfo[0], roomInfo[1], roomInfo[2])
			farm.Rooms = append(farm.Rooms, farm.Start)
			i++
			continue
		}
		if lines[i] == "##end" {
			roomInfo := strings.Split(lines[i+1], " ")
			farm.End = class.CreateRoom(roomInfo[0], roomInfo[1], roomInfo[2])
			farm.Rooms = append(farm.Rooms, farm.End)
			i++
			continue
		}
		// Récupération des rooms
		if roomInfo := strings.Split(lines[i], " "); len(roomInfo) == 3 {
			farm.Rooms = append(farm.Rooms, class.CreateRoom(roomInfo[0], roomInfo[1], roomInfo[2]))
		}
		// Récupération des liens
		if linkInfo := strings.Split(lines[i], "-"); len(linkInfo) == 2 {
			room1 := farm.GetRoomByName(linkInfo[0])
			room2 := farm.GetRoomByName(linkInfo[1])
			if room1 == nil || room2 == nil {
				return farm, errorsLem.ErrRoomLink
			}
			room1.RoomsLink = append(room1.RoomsLink, room2)
			room2.RoomsLink = append(room2.RoomsLink, room1)
		}
	}
	return farm, nil
}

package errorsLem

import "errors"

var (
	ErrNbLemin      = errors.New("invalid data format, invalid number of ants")
	ErrNoStartFound = errors.New("invalid data format, no start room found")
	ErrNoEndFound   = errors.New("invalid data format, no end room found")
	ErrNoPathFound  = errors.New("invalid data format, no path between start and end")
	ErrRoomLink     = errors.New("invalid data format, room does not exist in link")
	ErrNoArgs       = errors.New("go run . dataFileName.txt")
)

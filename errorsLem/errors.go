package errorsLem

import "errors"

var (
	ErrNbLemin      = errors.New("invalid data format, invalid number of Ants")
	ErrNoStartFound = errors.New("invalid data format, no start room found")
	ErrNoEndFound   = errors.New("invalid data format, no end room found")
	ErrNoPathFound  = errors.New("invalid data format, path between start and end")
	ErrRoomLink     = errors.New("invalid data format, room doest exist in link")
	ErrNoArgs       = errors.New("go run . dataFilleName.txt")
)

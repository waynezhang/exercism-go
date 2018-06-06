package queenattack

import (
	"errors"
	"math"
)

type position struct {
	x int
	y int
}

func parse(p string) (position, error) {
	if len(p) != 2 {
		return position{}, errors.New("")
	}

	var bytes = []byte(p)
	var pos = position{x: int(bytes[0]), y: int(bytes[1])}
	if 'a' <= pos.x && pos.x <= 'h' && '1' <= pos.y && pos.y <= '8' {
		return pos, nil
	}
	return pos, errors.New("")
}

// CanQueenAttack doc here
func CanQueenAttack(w, b string) (bool, error) {
	pos1, err := parse(w)
	if err != nil {
		return false, err
	}
	pos2, err := parse(b)
	if err != nil {
		return false, err
	}

	if pos1.x == pos2.x && pos1.y == pos2.y {
		return false, errors.New("")
	}

	canAttack := pos1.x == pos2.x || pos1.y == pos2.y || math.Abs(float64(pos1.x-pos2.x)) == math.Abs(float64(pos1.y-pos2.y))
	return canAttack, nil
}

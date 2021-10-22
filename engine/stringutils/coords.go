package stringutils

import (
	"errors"
	"fmt"
)

// ex: a1 -> 0, g8 -> 63
func CoordStringToIdx(coord *string) (uint16, error) {
	if len(*coord) != 2 {
		return 0, errors.New("Invalid coordinates " + *coord)
	}

	col := (*coord)[0] - 'a'
	row := (*coord)[1] - '1'
	if col < 0 || col > 7 || row < 0 || row > 7 {
		return 0, errors.New("Invalid coordinates " + *coord)
	}
	return uint16(row*8 + col), nil
}

// ex: 0 -> a1, 63 -> g8
func IdxToCoordString(idx uint16) *string {
	cs := fmt.Sprintf("%s%d", string(rune(idx % 8 + 'a')), idx / 8 + 1)
	return &cs
}
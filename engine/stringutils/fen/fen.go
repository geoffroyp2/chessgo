package fen

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/geoffroyp2/chessgo/engine/constants"
	"github.com/geoffroyp2/chessgo/engine/position"
	"github.com/geoffroyp2/chessgo/engine/stringutils"
)



func PositionFromFEN(FEN *string) *position.Position {

	newPos := position.NewPosition() // Empty position
	sections := strings.Fields(*FEN) 

	// Section 0: board description. Rows separated by '/'
	var BBIdx = map[rune]uint32{
		'K': constants.WHITEKING,
		'k': constants.BLACKKING,
		'Q': constants.WHITEQUEEN,
		'q': constants.BLACKQUEEN,
		'R': constants.WHITEROOK,
		'r': constants.BLACKROOK,
		'B': constants.WHITEBISHOP,
		'b': constants.BLACKBISHOP,
		'N': constants.WHITEKNIGHT,
		'n': constants.BLACKKNIGHT,
		'P': constants.WHITEPAWN,
		'p': constants.BLACKPAWN,
	}

	row := 7
	col := 0
	for _, char := range sections[0] {
		switch char {
		case '/': 
			col = 0
			row--
			break
		case 'K':	fallthrough
		case 'k':	fallthrough
		case 'Q':	fallthrough
		case 'q':	fallthrough
		case 'B':	fallthrough
		case 'b':	fallthrough
		case 'N':	fallthrough
		case 'n':	fallthrough
		case 'R':	fallthrough
		case 'r':	fallthrough
		case 'P':	fallthrough
		case 'p':
			newPos.Pieces[BBIdx[char]] |= 1 << (row * 8 + col)
			col++
			break
		case '1':	fallthrough
		case '2':	fallthrough
		case '3':	fallthrough
		case '4':	fallthrough
		case '5':	fallthrough
		case '6':	fallthrough
		case '7':	fallthrough
		case '8':
			charVal := int(char - '0')
			col += charVal
			break
		default:
			fmt.Println("Error: could not parse character ", string(char))
			break
		}
	}
	// Create union of all White & all Black pieces + occupied & empty sets
	for idx := 0; idx < 6; idx++ {
		newPos.AllWhite |= newPos.Pieces[idx]
	}
	for idx := 6; idx < 12; idx++ {
		newPos.AllBlack |= newPos.Pieces[idx]
	}
	newPos.Occupied = newPos.AllWhite | newPos.AllBlack
	newPos.Empty = ^newPos.Occupied

	// Section 1: player turn
	if sections[1] == "b" {
		newPos.SetPlayerTurn(1)
	}

	// Section 2: castle rights
	for idx, char := range sections[2] {
		if char != '-' {
			newPos.ToggleCastleRigth(uint16(idx))
		}
	}

	// Section 3: En-passant square
	val1, err := stringutils.CoordStringToIdx(&sections[3])
	if err != nil {
		newPos.SetEPSquare(constants.EPSQUARE_NONE)
	} else {
		newPos.SetEPSquare(val1)
	}
	

	// Section 4: Half-move Clock
	val2, err := strconv.Atoi(sections[4])
	if err != nil {
		fmt.Println("Could not parse Half-move counter: ", err)
	} else {
		newPos.HMClock = uint16(val2)
	}
	
	// Section 5: Move number // Not necessary
	// val2, err = strconv.Atoi(sections[5])
	// if err != nil {
	// 	fmt.Println("Could not parse Move Number: ", sections[5], err)
	// } else {
	// 	newPos.MoveNumber = uint32(val2)
	// }
	
	return newPos
}
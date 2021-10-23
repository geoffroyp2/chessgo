package position

import (
	"fmt"
	"strings"

	"github.com/geoffroyp2/chessgo/engine/constants"
	"github.com/geoffroyp2/chessgo/engine/stringutils"
)

func (pos *Position) PrintPosition() {
	var board [64]rune
	for i := 0; i < 64; i++ {
		board[i] = '.'
	}
	pieceRunes := [12]rune{'K', 'Q', 'R', 'B', 'N', 'P', 'k', 'q', 'r', 'b', 'n', 'p'}

	for pIdx := 0; pIdx < 12; pIdx++ {
		for offset := 0; offset < 64; offset++ {
			if (pos.Pieces[pIdx]>>offset)&1 != 0 {
				board[offset] = pieceRunes[pIdx]
			}
		}
	}

	var sb strings.Builder
	sb.WriteRune('\n')
	for row := 7; row >= 0; row-- {
		for col := 0; col < 8; col++ {
			sb.WriteRune(board[row * 8 + col])
			sb.WriteRune(' ')
		}
		switch row {
		case 6:
			sb.WriteString(fmt.Sprintf(" |  Next move: %s", *pos.getPlayerTurnString()))
			break
		case 5:
			sb.WriteString(fmt.Sprintf(" |  Castle: %s", *pos.getCastleRightsString()))
			break
		case 4:
			EPSquare := pos.GetEPSquare()
			EPSquareString := "-"
			if EPSquare != constants.EPSQUARE_NONE {
				EPSquareString = stringutils.IdxToCoordString(EPSquare)
			}
			sb.WriteString(fmt.Sprintf(" |  En-passant: %s", EPSquareString))
			break
		case 3: 
			sb.WriteString(fmt.Sprintf(" |  Half-moves clock: %d", pos.HMClock))
			break
		default:
			sb.WriteString(" |")
		}

		sb.WriteRune('\n')
	}

	fmt.Println(sb.String())
}

func (pos *Position) getCastleRightsString() *string {

	cr := pos.GetCastleRights()
	var sb strings.Builder
	runes := [4]rune{ 'q', 'k', 'Q', 'K' }

	for i := 3; i >= 0; i-- {
		if cr & (1 << i) != 0 {
			sb.WriteRune(runes[i])
		} else {
			sb.WriteRune('-')
		}
	}

	crString := sb.String()
	return &crString
}

func (pos *Position) getPlayerTurnString() *string {
	pt := pos.GetPlayerTurn()
	ptString := "b"

	if pt == 0 {
		ptString = "w"
	}

	return &ptString
}
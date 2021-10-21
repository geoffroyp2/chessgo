package movegen

import (
	"github.com/geoffroyp2/chessgo/engine/constants"
	"github.com/geoffroyp2/chessgo/engine/position"
)

func GenerateMoves(pos *position.Position, moveArray *constants.MoveArray) int {
	moveAmount := 0

	if pos.GetPlayerTurn() == 0 {

		moveAmount = wPawnMoves(pos, moveArray, moveAmount)

	} else {

		moveAmount = bPawnMoves(pos, moveArray, moveAmount)

	}

	return moveAmount
}
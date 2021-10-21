package movegen

import (
	"github.com/geoffroyp2/chessgo/engine/constants"
	"github.com/geoffroyp2/chessgo/engine/position"
)

func GenerateMoves(pos *position.Position, moveArray *constants.MoveArray) int {

	moveAmount := 0
	moveAmount = getPawnMoves(pos, moveArray, moveAmount)

	return moveAmount
}
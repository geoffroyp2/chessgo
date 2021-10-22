package movegen

import (
	"github.com/geoffroyp2/chessgo/engine/constants"
	"github.com/geoffroyp2/chessgo/engine/position"
)

func GenerateMoves(pos *position.Position, moveArray *constants.MoveArray) int {
	moveAmount := 0

	if pos.GetPlayerTurn() == 0 {

		moveAmount = wPawnMoves(pos, moveArray, moveAmount)
		moveAmount = wKnightMoves(pos, moveArray, moveAmount)
		moveAmount = wRookMoves(pos, moveArray, moveAmount)
		
	} else {
			
		moveAmount = bPawnMoves(pos, moveArray, moveAmount)
		moveAmount = bKnightMoves(pos, moveArray, moveAmount)
		moveAmount = bRookMoves(pos, moveArray, moveAmount)

	}

	return moveAmount
}

func getWCapturedPiece(pos *position.Position, idx uint32) uint32 {
	for pi := constants.WHITEKING; pi <= constants.WHITEPAWN; pi++ {
		if pos.Pieces[pi] & (1 << idx) != 0 {
			return pi
		}
	}
	return constants.NULLPIECE
}

func getBCapturedPiece(pos *position.Position, idx uint32) uint32 {
	for pi := constants.BLACKKING; pi <= constants.BLACKPAWN; pi++ {
		if pos.Pieces[pi] & (1 << idx) != 0 {
			return pi
		}
	}
	return constants.NULLPIECE
}
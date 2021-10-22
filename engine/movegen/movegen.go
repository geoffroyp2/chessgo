package movegen

import (
	"github.com/geoffroyp2/chessgo/engine/constants"
	"github.com/geoffroyp2/chessgo/engine/position"
)

func GenerateMoves(pos *position.Position, moveArray *constants.MoveArray) int {
	moveAmount := 0

	if pos.GetPlayerTurn() == constants.PLAYERWHITE {

		ennemyAttacks := bQueenAttacks(pos)
		ennemyAttacks |= bKingAttacks(pos)
		ennemyAttacks |= bRookAttacks(pos)
		ennemyAttacks |= bBishopAttacks(pos)
		ennemyAttacks |= bKnightAttacks(pos)
		ennemyAttacks |= bPawnAttacks(pos)

		moveAmount = wKingMoves(pos, moveArray, moveAmount, ^ennemyAttacks)
		moveAmount = wQueenMoves(pos, moveArray, moveAmount)
		moveAmount = wRookMoves(pos, moveArray, moveAmount)
		moveAmount = wBishopMoves(pos, moveArray, moveAmount)
		moveAmount = wKnightMoves(pos, moveArray, moveAmount)
		moveAmount = wPawnMoves(pos, moveArray, moveAmount)
		
	} else {

		ennemyAttacks := wQueenAttacks(pos)
		ennemyAttacks |= wKingAttacks(pos)
		ennemyAttacks |= wRookAttacks(pos)
		ennemyAttacks |= wBishopAttacks(pos)
		ennemyAttacks |= wKnightAttacks(pos)
		ennemyAttacks |= wPawnAttacks(pos)
		
		moveAmount = bKingMoves(pos, moveArray, moveAmount, ^ennemyAttacks)
		moveAmount = bQueenMoves(pos, moveArray, moveAmount)
		moveAmount = bRookMoves(pos, moveArray, moveAmount)
		moveAmount = bBishopMoves(pos, moveArray, moveAmount)
		moveAmount = bKnightMoves(pos, moveArray, moveAmount)
		moveAmount = bPawnMoves(pos, moveArray, moveAmount)

	}

	return moveAmount
}

func getWCapturedPiece(pos *position.Position, idx uint32) uint32 {
	for pi := constants.WHITEKING; pi <= constants.WHITEPAWN; pi++ {
		if pos.Pieces[pi] & (1<<idx) != 0 {
			return pi
		}
	}
	return constants.NULLPIECE
}

func getBCapturedPiece(pos *position.Position, idx uint32) uint32 {
	for pi := constants.BLACKKING; pi <= constants.BLACKPAWN; pi++ {
		if pos.Pieces[pi] & (1<<idx) != 0 {
			return pi
		}
	}
	return constants.NULLPIECE
}
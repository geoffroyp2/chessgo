package movegen

import (
	"github.com/geoffroyp2/chessgo/engine/constants"
	"github.com/geoffroyp2/chessgo/engine/makemove"
	"github.com/geoffroyp2/chessgo/engine/position"
)

func GenerateMoves(pos *position.Position, moveArray *constants.MoveArray) int {
	moveAmount := 0

	if pos.GetPlayerTurn() == constants.PLAYERWHITE {

		moveAmount = wQueenMoves(pos, moveArray, moveAmount)
		moveAmount = wRookMoves(pos, moveArray, moveAmount)
		moveAmount = wBishopMoves(pos, moveArray, moveAmount)
		moveAmount = wKnightMoves(pos, moveArray, moveAmount)
		moveAmount = wPawnMoves(pos, moveArray, moveAmount)

		// Check validity of moves and remove invalid ones
		for i := 0; i < moveAmount; i++ {
			posCopy := position.CopyPosition(pos)
			makemove.MakeMove(posCopy, moveArray[i])
			ennemyAttacks := getBAttacks(posCopy)
			if ennemyAttacks & posCopy.Pieces[constants.WHITEKING] != 0 {
				// move.PrintMove(moveArray[i])
				moveArray[i] = moveArray[moveAmount - 1]
				moveAmount--
				i--
			}
		}
		
		ennemyAttacks := getBAttacks(pos)
		moveAmount = wKingMoves(pos, moveArray, moveAmount, ^ennemyAttacks)
		
	} else {

		moveAmount = bQueenMoves(pos, moveArray, moveAmount)
		moveAmount = bRookMoves(pos, moveArray, moveAmount)
		moveAmount = bBishopMoves(pos, moveArray, moveAmount)
		moveAmount = bKnightMoves(pos, moveArray, moveAmount)
		moveAmount = bPawnMoves(pos, moveArray, moveAmount)

		for i := 0; i < moveAmount; i++ {
			posCopy := position.CopyPosition(pos)
			makemove.MakeMove(posCopy, moveArray[i])
			ennemyAttacks := getWAttacks(posCopy)
			if ennemyAttacks & posCopy.Pieces[constants.BLACKKING] != 0 {
				moveArray[i] = moveArray[moveAmount - 1]
				moveAmount--
				i--
			}
		}
		
		ennemyAttacks := getWAttacks(pos)
		moveAmount = bKingMoves(pos, moveArray, moveAmount, ^ennemyAttacks)
		
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

func getWAttacks(pos *position.Position) uint64 {
	ennemyAttacks := wQueenAttacks(pos)
	ennemyAttacks |= wKingAttacks(pos)
	ennemyAttacks |= wRookAttacks(pos)
	ennemyAttacks |= wBishopAttacks(pos)
	ennemyAttacks |= wKnightAttacks(pos)
	return ennemyAttacks | wPawnAttacks(pos)
}

func getBAttacks(pos *position.Position) uint64 {
	ennemyAttacks := bQueenAttacks(pos)
	ennemyAttacks |= bKingAttacks(pos)
	ennemyAttacks |= bRookAttacks(pos)
	ennemyAttacks |= bBishopAttacks(pos)
	ennemyAttacks |= bKnightAttacks(pos)
	return ennemyAttacks | bPawnAttacks(pos)
}


package movegen

/*
https://www.chessprogramming.org/Sliding_Piece_Attacks
https://www.chessprogramming.org/Classical_Approach
*/

import (
	"math/bits"

	"github.com/geoffroyp2/chessgo/engine/constants"
	"github.com/geoffroyp2/chessgo/engine/position"
)

func wRookMoves(pos *position.Position, moveArray *constants.MoveArray, moveAmount int) int {

	rooks := pos.Pieces[constants.WHITEROOK]

	for rooks != 0 {
		rookIdx := uint32(bits.TrailingZeros64(rooks))

		moveAmount = processWPosRay(rookIdx, constants.DIRN, constants.WHITEROOK, pos, moveArray, moveAmount)
		moveAmount = processWPosRay(rookIdx, constants.DIRE, constants.WHITEROOK, pos, moveArray, moveAmount)
		moveAmount = processWNegRay(rookIdx, constants.DIRS, constants.WHITEROOK, pos, moveArray, moveAmount)
		moveAmount = processWNegRay(rookIdx, constants.DIRW, constants.WHITEROOK, pos, moveArray, moveAmount)

		rooks ^= 1<<rookIdx
	}

	return moveAmount
}

func bRookMoves(pos *position.Position, moveArray *constants.MoveArray, moveAmount int) int {

	rooks := pos.Pieces[constants.BLACKROOK]

	for rooks != 0 {
		rookIdx := uint32(bits.TrailingZeros64(rooks))

		moveAmount = processBPosRay(rookIdx, constants.DIRN, constants.BLACKROOK, pos, moveArray, moveAmount)
		moveAmount = processBPosRay(rookIdx, constants.DIRE, constants.BLACKROOK, pos, moveArray, moveAmount)
		moveAmount = processBNegRay(rookIdx, constants.DIRS, constants.BLACKROOK, pos, moveArray, moveAmount)
		moveAmount = processBNegRay(rookIdx, constants.DIRW, constants.BLACKROOK, pos, moveArray, moveAmount)

		rooks ^= 1<<rookIdx
	}

	return moveAmount
}
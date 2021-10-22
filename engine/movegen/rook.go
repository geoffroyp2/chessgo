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

// Generate a simple BitBoard of squares that can be attacked by rooks
func wRookAttacks(pos *position.Position) uint64 {
	var attacks uint64 = 0
	rooks := pos.Pieces[constants.WHITEROOK]
	for rooks != 0 {
		rookIdx := bits.TrailingZeros64(rooks)
		ray1, _ := getPosRay(pos.Occupied, uint32(rookIdx), constants.DIRN)
		ray2, _ := getPosRay(pos.Occupied, uint32(rookIdx), constants.DIRE)
		ray3, _ := getNegRay(pos.Occupied, uint32(rookIdx), constants.DIRS)
		ray4, _ := getNegRay(pos.Occupied, uint32(rookIdx), constants.DIRW)
		attacks |= ray1 | ray2 | ray3 | ray4
		rooks ^= 1<<rookIdx
	}
	return attacks
}

func bRookAttacks(pos *position.Position) uint64 {
	var attacks uint64 = 0
	rooks := pos.Pieces[constants.BLACKROOK]
	for rooks != 0 {
		rookIdx := bits.TrailingZeros64(rooks)
		ray1, _ := getPosRay(pos.Occupied, uint32(rookIdx), constants.DIRN)
		ray2, _ := getPosRay(pos.Occupied, uint32(rookIdx), constants.DIRE)
		ray3, _ := getNegRay(pos.Occupied, uint32(rookIdx), constants.DIRS)
		ray4, _ := getNegRay(pos.Occupied, uint32(rookIdx), constants.DIRW)
		attacks |= ray1 | ray2 | ray3 | ray4
		rooks ^= 1<<rookIdx
	}
	return attacks
}
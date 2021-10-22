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

func wBishopMoves(pos *position.Position, moveArray *constants.MoveArray, moveAmount int) int {

	bishops := pos.Pieces[constants.WHITEBISHOP]

	for bishops != 0 {
		bishopIdx := uint32(bits.TrailingZeros64(bishops))

		moveAmount = processWPosRay(bishopIdx, constants.DIRNE, constants.WHITEBISHOP, pos, moveArray, moveAmount)
		moveAmount = processWPosRay(bishopIdx, constants.DIRNW, constants.WHITEBISHOP, pos, moveArray, moveAmount)
		moveAmount = processWNegRay(bishopIdx, constants.DIRSE, constants.WHITEBISHOP, pos, moveArray, moveAmount)
		moveAmount = processWNegRay(bishopIdx, constants.DIRSW, constants.WHITEBISHOP, pos, moveArray, moveAmount)

		bishops ^= 1<<bishopIdx
	}

	return moveAmount
}

func bBishopMoves(pos *position.Position, moveArray *constants.MoveArray, moveAmount int) int {

	bishops := pos.Pieces[constants.BLACKBISHOP]

	for bishops != 0 {
		bishopIdx := uint32(bits.TrailingZeros64(bishops))

		moveAmount = processBPosRay(bishopIdx, constants.DIRNE, constants.BLACKBISHOP, pos, moveArray, moveAmount)
		moveAmount = processBPosRay(bishopIdx, constants.DIRNW, constants.BLACKBISHOP, pos, moveArray, moveAmount)
		moveAmount = processBNegRay(bishopIdx, constants.DIRSE, constants.BLACKBISHOP, pos, moveArray, moveAmount)
		moveAmount = processBNegRay(bishopIdx, constants.DIRSW, constants.BLACKBISHOP, pos, moveArray, moveAmount)

		bishops ^= 1<<bishopIdx
	}

	return moveAmount
}

// Generate a simple BitBoard of squares that can be attacked by bishops
func wBishopAttacks(pos *position.Position) uint64 {
	var attacks uint64 = 0
	bishops := pos.Pieces[constants.WHITEBISHOP]
	for bishops != 0 {
		bishopIdx := bits.TrailingZeros64(bishops)
		ray, _ := getPosRay(pos.Occupied, uint32(bishopIdx), constants.DIRNE)
		attacks |= ray
		ray, _  = getPosRay(pos.Occupied, uint32(bishopIdx), constants.DIRNW)
		attacks |= ray
		ray, _  = getNegRay(pos.Occupied, uint32(bishopIdx), constants.DIRSE)
		attacks |= ray
		ray, _  = getNegRay(pos.Occupied, uint32(bishopIdx), constants.DIRSW)
		attacks |= ray
		bishops ^= 1<<bishopIdx
	}
	return attacks
}

func bBishopAttacks(pos *position.Position) uint64 {
	var attacks uint64 = 0
	bishops := pos.Pieces[constants.BLACKBISHOP]
	for bishops != 0 {
		bishopIdx := bits.TrailingZeros64(bishops)
		ray, _ := getPosRay(pos.Occupied, uint32(bishopIdx), constants.DIRNE)
		attacks |= ray
		ray, _  = getPosRay(pos.Occupied, uint32(bishopIdx), constants.DIRNW)
		attacks |= ray
		ray, _  = getNegRay(pos.Occupied, uint32(bishopIdx), constants.DIRSE)
		attacks |= ray
		ray, _  = getNegRay(pos.Occupied, uint32(bishopIdx), constants.DIRSW)
		attacks |= ray
		bishops ^= 1<<bishopIdx
	}
	return attacks
}
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

func wQueenMoves(pos *position.Position, moveArray *constants.MoveArray, moveAmount int) int {

	queens := pos.Pieces[constants.WHITEQUEEN]

	for queens != 0 {
		queenIdx := uint32(bits.TrailingZeros64(queens))

		moveAmount = processWPosRay(queenIdx, constants.DIRN, constants.WHITEQUEEN, pos, moveArray, moveAmount)
		moveAmount = processWPosRay(queenIdx, constants.DIRE, constants.WHITEQUEEN, pos, moveArray, moveAmount)
		moveAmount = processWNegRay(queenIdx, constants.DIRS, constants.WHITEQUEEN, pos, moveArray, moveAmount)
		moveAmount = processWNegRay(queenIdx, constants.DIRW, constants.WHITEQUEEN, pos, moveArray, moveAmount)

		moveAmount = processWPosRay(queenIdx, constants.DIRNE, constants.WHITEQUEEN, pos, moveArray, moveAmount)
		moveAmount = processWPosRay(queenIdx, constants.DIRNW, constants.WHITEQUEEN, pos, moveArray, moveAmount)
		moveAmount = processWNegRay(queenIdx, constants.DIRSE, constants.WHITEQUEEN, pos, moveArray, moveAmount)
		moveAmount = processWNegRay(queenIdx, constants.DIRSW, constants.WHITEQUEEN, pos, moveArray, moveAmount)

		queens ^= 1<<queenIdx
	}

	return moveAmount
}

func bQueenMoves(pos *position.Position, moveArray *constants.MoveArray, moveAmount int) int {

	queens := pos.Pieces[constants.BLACKQUEEN]

	for queens != 0 {
		queenIdx := uint32(bits.TrailingZeros64(queens))

		moveAmount = processBPosRay(queenIdx, constants.DIRN, constants.BLACKQUEEN, pos, moveArray, moveAmount)
		moveAmount = processBPosRay(queenIdx, constants.DIRE, constants.BLACKQUEEN, pos, moveArray, moveAmount)
		moveAmount = processBNegRay(queenIdx, constants.DIRS, constants.BLACKQUEEN, pos, moveArray, moveAmount)
		moveAmount = processBNegRay(queenIdx, constants.DIRW, constants.BLACKQUEEN, pos, moveArray, moveAmount)

		moveAmount = processBPosRay(queenIdx, constants.DIRNE, constants.BLACKQUEEN, pos, moveArray, moveAmount)
		moveAmount = processBPosRay(queenIdx, constants.DIRNW, constants.BLACKQUEEN, pos, moveArray, moveAmount)
		moveAmount = processBNegRay(queenIdx, constants.DIRSE, constants.BLACKQUEEN, pos, moveArray, moveAmount)
		moveAmount = processBNegRay(queenIdx, constants.DIRSW, constants.BLACKQUEEN, pos, moveArray, moveAmount)

		queens ^= 1<<queenIdx
	}

	return moveAmount
}

// Generate a simple BitBoard of squares that can be attacked by queens
func wQueenAttacks(pos *position.Position) uint64 {
	var attacks uint64 = 0
	queens := pos.Pieces[constants.WHITEQUEEN]
	for queens != 0 {
		queenIdx := bits.TrailingZeros64(queens)
		ray1, _ := getPosRay(pos.Occupied, uint32(queenIdx), constants.DIRN)
		ray2, _ := getPosRay(pos.Occupied, uint32(queenIdx), constants.DIRE)
		ray3, _ := getNegRay(pos.Occupied, uint32(queenIdx), constants.DIRS)
		ray4, _ := getNegRay(pos.Occupied, uint32(queenIdx), constants.DIRW)
		ray5, _ := getPosRay(pos.Occupied, uint32(queenIdx), constants.DIRNE)
		ray6, _ := getPosRay(pos.Occupied, uint32(queenIdx), constants.DIRNW)
		ray7, _ := getNegRay(pos.Occupied, uint32(queenIdx), constants.DIRSE)
		ray8, _ := getNegRay(pos.Occupied, uint32(queenIdx), constants.DIRSW)
		attacks |= ray1 | ray2 | ray3 | ray4 | ray5 | ray6 | ray7 | ray8
		queens ^= 1<<queenIdx
	}
	return attacks
}

func bQueenAttacks(pos *position.Position) uint64 {
	var attacks uint64 = 0
	queens := pos.Pieces[constants.BLACKQUEEN]
	for queens != 0 {
		queenIdx := bits.TrailingZeros64(queens)
		ray1, _ := getPosRay(pos.Occupied, uint32(queenIdx), constants.DIRN)
		ray2, _ := getPosRay(pos.Occupied, uint32(queenIdx), constants.DIRE)
		ray3, _ := getNegRay(pos.Occupied, uint32(queenIdx), constants.DIRS)
		ray4, _ := getNegRay(pos.Occupied, uint32(queenIdx), constants.DIRW)
		ray5, _ := getPosRay(pos.Occupied, uint32(queenIdx), constants.DIRNE)
		ray6, _ := getPosRay(pos.Occupied, uint32(queenIdx), constants.DIRNW)
		ray7, _ := getNegRay(pos.Occupied, uint32(queenIdx), constants.DIRSE)
		ray8, _ := getNegRay(pos.Occupied, uint32(queenIdx), constants.DIRSW)
		attacks |= ray1 | ray2 | ray3 | ray4 | ray5 | ray6 | ray7 | ray8
		queens ^= 1<<queenIdx
	}
	return attacks
}
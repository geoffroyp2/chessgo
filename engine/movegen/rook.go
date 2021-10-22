package movegen

/*
https://www.chessprogramming.org/Sliding_Piece_Attacks
https://www.chessprogramming.org/Classical_Approach
*/

import (
	"math/bits"

	"github.com/geoffroyp2/chessgo/engine/constants"
	"github.com/geoffroyp2/chessgo/engine/move"
	"github.com/geoffroyp2/chessgo/engine/position"
)

func wRookMoves(pos *position.Position, moveArray *constants.MoveArray, moveAmount int) int {

	rooks := pos.Pieces[constants.WHITEROOK]

	for rooks != 0 {
		rookIdx := uint32(bits.TrailingZeros64(rooks))

		attacksN, isBlockedN := getPosRay(pos.Occupied, rookIdx, constants.DIRN)
		if isBlockedN {
			destIdx := uint32(bits.Len64(attacksN) - 1)
			if pos.AllBlack & (1<<destIdx) != 0 {
				capturedPiece := getBCapturedPiece(pos, destIdx)
				(*moveArray)[moveAmount] = move.CreateMove(rookIdx, destIdx, move.CAPTURE, constants.WHITEROOK, capturedPiece)
				moveAmount++
			}
			attacksN ^= 1<<destIdx
		}
		for attacksN != 0 {
			destIdx := uint32(bits.TrailingZeros64(attacksN))
			(*moveArray)[moveAmount] = move.CreateMove(rookIdx, destIdx, move.MOVE, constants.WHITEROOK, constants.NULLPIECE)
			moveAmount++
			attacksN ^= 1<<destIdx
		}

		attacksE, isBlockedE := getPosRay(pos.Occupied, rookIdx, constants.DIRE)
		if isBlockedE {
			destIdx := uint32(bits.Len64(attacksE) - 1)
			if pos.AllBlack & (1<<destIdx) != 0 {
				capturedPiece := getBCapturedPiece(pos, destIdx)
				(*moveArray)[moveAmount] = move.CreateMove(rookIdx, destIdx, move.CAPTURE, constants.WHITEROOK, capturedPiece)
				moveAmount++
			}
			attacksE ^= 1<<destIdx
		}
		for attacksE != 0 {
			destIdx := uint32(bits.TrailingZeros64(attacksE))
			(*moveArray)[moveAmount] = move.CreateMove(rookIdx, destIdx, move.MOVE, constants.WHITEROOK, constants.NULLPIECE)
			moveAmount++
			attacksE ^= 1<<destIdx
		}

		attacksS, isBlockedS := getNegRay(pos.Occupied, rookIdx, constants.DIRS)
		if isBlockedS {
			destIdx := uint32(bits.TrailingZeros64(attacksS))
			if pos.AllBlack & (1<<destIdx) != 0 {
				capturedPiece := getBCapturedPiece(pos, destIdx)
				(*moveArray)[moveAmount] = move.CreateMove(rookIdx, destIdx, move.CAPTURE, constants.WHITEROOK, capturedPiece)
				moveAmount++
			}
			attacksS ^= 1<<destIdx
		}
		for attacksS != 0 {
			destIdx := uint32(bits.Len64(attacksS) - 1)
			(*moveArray)[moveAmount] = move.CreateMove(rookIdx, destIdx, move.MOVE, constants.WHITEROOK, constants.NULLPIECE)
			moveAmount++
			attacksS ^= 1<<destIdx
		}

		attacksW, isBlockedW := getNegRay(pos.Occupied, rookIdx, constants.DIRW)
		if isBlockedW {
			destIdx := uint32(bits.TrailingZeros64(attacksW))
			if pos.AllBlack & (1<<destIdx) != 0 {
				capturedPiece := getBCapturedPiece(pos, destIdx)
				(*moveArray)[moveAmount] = move.CreateMove(rookIdx, destIdx, move.CAPTURE, constants.WHITEROOK, capturedPiece)
				moveAmount++
			}
			attacksW ^= 1<<destIdx
		}
		for attacksW != 0 {
			destIdx := uint32(bits.Len64(attacksW) - 1)
			(*moveArray)[moveAmount] = move.CreateMove(rookIdx, destIdx, move.MOVE, constants.WHITEROOK, constants.NULLPIECE)
			moveAmount++
			attacksW ^= 1<<destIdx
		}

		rooks ^= 1<<rookIdx
	}

	return moveAmount
}

func bRookMoves(pos *position.Position, moveArray *constants.MoveArray, moveAmount int) int {

	rooks := pos.Pieces[constants.BLACKROOK]

	for rooks != 0 {
		rookIdx := uint32(bits.TrailingZeros64(rooks))

		attacksN, isBlockedN := getPosRay(pos.Occupied, rookIdx, constants.DIRN)
		if isBlockedN {
			destIdx := uint32(bits.Len64(attacksN) - 1)
			if pos.AllWhite & (1<<destIdx) != 0 {
				capturedPiece := getWCapturedPiece(pos, destIdx)
				(*moveArray)[moveAmount] = move.CreateMove(rookIdx, destIdx, move.CAPTURE, constants.BLACKROOK, capturedPiece)
				moveAmount++
			}
			attacksN ^= 1<<destIdx
		}
		for attacksN != 0 {
			destIdx := uint32(bits.TrailingZeros64(attacksN))
			(*moveArray)[moveAmount] = move.CreateMove(rookIdx, destIdx, move.MOVE, constants.BLACKROOK, constants.NULLPIECE)
			moveAmount++
			attacksN ^= 1<<destIdx
		}

		attacksE, isBlockedE := getPosRay(pos.Occupied, rookIdx, constants.DIRE)
		if isBlockedE {
			destIdx := uint32(bits.Len64(attacksE) - 1)
			if pos.AllWhite & (1<<destIdx) != 0 {
				capturedPiece := getWCapturedPiece(pos, destIdx)
				(*moveArray)[moveAmount] = move.CreateMove(rookIdx, destIdx, move.CAPTURE, constants.BLACKROOK, capturedPiece)
				moveAmount++
			}
			attacksE ^= 1<<destIdx
		}
		for attacksE != 0 {
			destIdx := uint32(bits.TrailingZeros64(attacksE))
			(*moveArray)[moveAmount] = move.CreateMove(rookIdx, destIdx, move.MOVE, constants.BLACKROOK, constants.NULLPIECE)
			moveAmount++
			attacksE ^= 1<<destIdx
		}

		attacksS, isBlockedS := getNegRay(pos.Occupied, rookIdx, constants.DIRS)
		if isBlockedS {
			destIdx := uint32(bits.TrailingZeros64(attacksS))
			if pos.AllWhite & (1<<destIdx) != 0 {
				capturedPiece := getWCapturedPiece(pos, destIdx)
				(*moveArray)[moveAmount] = move.CreateMove(rookIdx, destIdx, move.CAPTURE, constants.BLACKROOK, capturedPiece)
				moveAmount++
			}
			attacksS ^= 1<<destIdx
		}
		for attacksS != 0 {
			destIdx := uint32(bits.Len64(attacksS) - 1)
			(*moveArray)[moveAmount] = move.CreateMove(rookIdx, destIdx, move.MOVE, constants.BLACKROOK, constants.NULLPIECE)
			moveAmount++
			attacksS ^= 1<<destIdx
		}

		attacksW, isBlockedW := getNegRay(pos.Occupied, rookIdx, constants.DIRW)
		if isBlockedW {
			destIdx := uint32(bits.TrailingZeros64(attacksW))
			if pos.AllWhite & (1<<destIdx) != 0 {
				capturedPiece := getWCapturedPiece(pos, destIdx)
				(*moveArray)[moveAmount] = move.CreateMove(rookIdx, destIdx, move.CAPTURE, constants.BLACKROOK, capturedPiece)
				moveAmount++
			}
			attacksW ^= 1<<destIdx
		}
		for attacksW != 0 {
			destIdx := uint32(bits.Len64(attacksW) - 1)
			(*moveArray)[moveAmount] = move.CreateMove(rookIdx, destIdx, move.MOVE, constants.BLACKROOK, constants.NULLPIECE)
			moveAmount++
			attacksW ^= 1<<destIdx
		}

		rooks ^= 1<<rookIdx
	}

	return moveAmount
}
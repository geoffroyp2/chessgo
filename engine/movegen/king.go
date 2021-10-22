package movegen

import (
	"math/bits"

	"github.com/geoffroyp2/chessgo/engine/constants"
	"github.com/geoffroyp2/chessgo/engine/move"
	"github.com/geoffroyp2/chessgo/engine/position"
)

var kingMovesLookup = [64]uint64{
	0x302, 0x705, 0xE0A, 0x1C14, 0x3828, 0x7050, 0xE0A0, 0xC040,
	0x30203  << (8*0), 0x70507  << (8*0), 0xE0A0E  << (8*0), 0x1C141C << (8*0),
	0x382838 << (8*0), 0x705070 << (8*0), 0xE0A0E0 << (8*0), 0xC040C0 << (8*0),
	0x30203  << (8*1), 0x70507  << (8*1), 0xE0A0E  << (8*1), 0x1C141C << (8*1),
	0x382838 << (8*1), 0x705070 << (8*1), 0xE0A0E0 << (8*1), 0xC040C0 << (8*1),
	0x30203  << (8*2), 0x70507  << (8*2), 0xE0A0E  << (8*2), 0x1C141C << (8*2),
	0x382838 << (8*2), 0x705070 << (8*2), 0xE0A0E0 << (8*2), 0xC040C0 << (8*2),
	0x30203  << (8*3), 0x70507  << (8*3), 0xE0A0E  << (8*3), 0x1C141C << (8*3),
	0x382838 << (8*3), 0x705070 << (8*3), 0xE0A0E0 << (8*3), 0xC040C0 << (8*3),
	0x30203  << (8*4), 0x70507  << (8*4), 0xE0A0E  << (8*4), 0x1C141C << (8*4),
	0x382838 << (8*4), 0x705070 << (8*4), 0xE0A0E0 << (8*4), 0xC040C0 << (8*4),
	0x30203  << (8*5), 0x70507  << (8*5), 0xE0A0E  << (8*5), 0x1C141C << (8*5),
	0x382838 << (8*5), 0x705070 << (8*5), 0xE0A0E0 << (8*5), 0xC040C0 << (8*5),
	0x203    << (8*6), 0x507    << (8*6), 0xA0E    << (8*6), 0x141C   << (8*6), 
	0x2838   << (8*6), 0x5070   << (8*6), 0xA0E0   << (8*6), 0x40C0   << (8*6),
}

func wKingMoves(pos *position.Position, moveArray *constants.MoveArray, moveAmount int, safeFromEnnemy uint64) int {

	king := pos.Pieces[constants.WHITEKING]
	kingIdx := uint32(bits.TrailingZeros64(king))
	destinations := kingMovesLookup[kingIdx] & pos.Empty & safeFromEnnemy
	captures := kingMovesLookup[kingIdx] & pos.AllBlack & safeFromEnnemy

	for destinations != 0 {
		destIdx := uint32(bits.TrailingZeros64(destinations))
		(*moveArray)[moveAmount] = move.CreateMove(kingIdx, destIdx, move.MOVE, constants.WHITEKING, constants.NULLPIECE)
		moveAmount++
		destinations ^= 1 << destIdx
	}

	for captures != 0 {
		destIdx := uint32(bits.TrailingZeros64(captures))
		capturedPiece := getBCapturedPiece(pos, destIdx)
		(*moveArray)[moveAmount] = move.CreateMove(kingIdx, destIdx, move.CAPTURE, constants.WHITEKING, capturedPiece)
		moveAmount++
		captures ^= 1 << destIdx
	}

	return moveAmount
}

func bKingMoves(pos *position.Position, moveArray *constants.MoveArray, moveAmount int, safeFromEnnemy uint64) int {

	king := pos.Pieces[constants.BLACKKING]
	kingIdx := uint32(bits.TrailingZeros64(king))
	destinations := kingMovesLookup[kingIdx] & pos.Empty & safeFromEnnemy
	captures := kingMovesLookup[kingIdx] & pos.AllWhite & safeFromEnnemy

	for destinations != 0 {
		destIdx := uint32(bits.TrailingZeros64(destinations))
		(*moveArray)[moveAmount] = move.CreateMove(kingIdx, destIdx, move.MOVE, constants.BLACKKING, constants.NULLPIECE)
		moveAmount++
		destinations ^= 1 << destIdx
	}

	for captures != 0 {
		destIdx := uint32(bits.TrailingZeros64(captures))
		capturedPiece := getBCapturedPiece(pos, destIdx)
		(*moveArray)[moveAmount] = move.CreateMove(kingIdx, destIdx, move.CAPTURE, constants.BLACKKING, capturedPiece)
		moveAmount++
		captures ^= 1 << destIdx
	}

	return moveAmount
}

// Generate a simple BitBoard of squares that can be attacked by the king
func wKingAttacks(pos *position.Position) uint64 {
	kingIdx := bits.TrailingZeros64(pos.Pieces[constants.WHITEKNIGHT])
	return kingMovesLookup[kingIdx]
}

func bKingAttacks(pos *position.Position) uint64 {
	kingIdx := bits.TrailingZeros64(pos.Pieces[constants.BLACKKING])
	return kingMovesLookup[kingIdx]
}
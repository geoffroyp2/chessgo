package movegen

/*
https://www.chessprogramming.org/Knight_Pattern
*/

import (
	"math/bits"

	"github.com/geoffroyp2/chessgo/engine/constants"
	"github.com/geoffroyp2/chessgo/engine/move"
	"github.com/geoffroyp2/chessgo/engine/position"
)

// Patterns for knight moves indexed by starting square
var knightMovesLookup = [64]uint64{
	0x20400,            0x50800,            0xA1100,            0x142200,           
	0x284400,           0x508800,           0xA01000,           0x402000,
	0x2040004,          0x5080008,          0xA110011,          0x14220022,         
	0x28440044,         0x50880088,         0xA0100010,         0x40200020,
	0x204000402,        0x508000805,        0xA1100110A,        0x1422002214,       
	0x2844004428,       0x5088008850,       0xA0100010A0,       0x4020002040, 
	0x20400040200,      0x50800080500,      0xA1100110A00,      0x142200221400,     
	0x284400442800,     0x508800885000,     0xA0100010A000,     0x402000204000,
	0x2040004020000,    0x5080008050000,    0xA1100110A0000,    0x14220022140000,   
	0x28440044280000,   0x50880088500000,   0xA0100010A00000,   0x40200020400000,
	0x204000402000000,  0x508000805000000,  0xA1100110A000000,  0x1422002214000000, 
	0x2844004428000000, 0x5088008850000000, 0xA0100010A0000000, 0x4020002040000000,
	0x400040200000000,  0x800080500000000,  0x1100110A00000000, 0x2200221400000000, 
	0x4400442800000000, 0x8800885000000000, 0x100010A000000000, 0x2000204000000000,
	0x4020000000000,    0x8050000000000,    0x110A0000000000,   0x22140000000000,   
	0x44280000000000,   0x88500000000000,   0x10A00000000000,   0x20400000000000,
}

func wKnightMoves(pos *position.Position, moveArray *constants.MoveArray, moveAmount int) int {

	knights := pos.Pieces[constants.WHITEKNIGHT]

	for knights != 0 {
		knightIdx := uint32(bits.TrailingZeros64(knights))
		destinations := knightMovesLookup[knightIdx] & pos.Empty
		captures := knightMovesLookup[knightIdx] & pos.AllBlack

		for destinations != 0 {
			destIdx := uint32(bits.TrailingZeros64(destinations))
			(*moveArray)[moveAmount] = move.CreateMove(knightIdx, destIdx, move.MOVE, constants.WHITEKNIGHT, constants.NULLPIECE)
			moveAmount++
			destinations ^= 1 << destIdx
		}

		for captures != 0 {

			destIdx := uint32(bits.TrailingZeros64(captures))
			capturedPiece := constants.NULLPIECE
			for pi := constants.BLACKKING; pi <= constants.BLACKPAWN; pi++ {
				if pos.Pieces[pi] & (1 << destIdx) != 0 {
					capturedPiece = pi
				}
			}

			(*moveArray)[moveAmount] = move.CreateMove(knightIdx, destIdx, move.CAPTURE, constants.WHITEKNIGHT, capturedPiece)
			moveAmount++
			captures ^= 1 << destIdx
		}

		knights ^= 1 << knightIdx
	}

	return moveAmount
}

func bKnightMoves(pos *position.Position, moveArray *constants.MoveArray, moveAmount int) int {

	knights := pos.Pieces[constants.BLACKKNIGHT]

	for knights != 0 {
		knightIdx := uint32(bits.TrailingZeros64(knights))
		destinations := knightMovesLookup[knightIdx] & pos.Empty
		captures := knightMovesLookup[knightIdx] & pos.AllWhite

		for destinations != 0 {
			destIdx := uint32(bits.TrailingZeros64(destinations))
			(*moveArray)[moveAmount] = move.CreateMove(knightIdx, destIdx, move.MOVE, constants.BLACKKNIGHT, constants.NULLPIECE)
			moveAmount++
			destinations ^= 1 << destIdx
		}

		for captures != 0 {

			destIdx := uint32(bits.TrailingZeros64(captures))
			capturedPiece := constants.NULLPIECE
			for pi := constants.WHITEKING; pi <= constants.WHITEPAWN; pi++ {
				if pos.Pieces[pi] & (1 << destIdx) != 0 {
					capturedPiece = pi
				}
			}

			(*moveArray)[moveAmount] = move.CreateMove(knightIdx, destIdx, move.CAPTURE, constants.BLACKKNIGHT, capturedPiece)
			moveAmount++
			captures ^= 1 << destIdx
		}

		knights ^= 1 << knightIdx
	}

	return moveAmount
}

// Generate a simple BitBoard of squares that can be attacked by knights
func wKnightAttacks(pos *position.Position) uint64 {
	var attacks uint64 = 0
	knights := pos.Pieces[constants.WHITEKNIGHT]
	for knights != 0 {
		knightIdx := bits.TrailingZeros64(knights)
		attacks |= knightMovesLookup[knightIdx]
		knights ^= 1 << knightIdx
	}
	return attacks
}

func bKnightAttacks(pos *position.Position) uint64 {
	var attacks uint64 = 0
	knights := pos.Pieces[constants.BLACKKNIGHT]
	for knights != 0 {
		knightIdx := bits.TrailingZeros64(knights)
		attacks |= knightMovesLookup[knightIdx]
		knights ^= 1 << knightIdx
	}
	return attacks
}
package movegen

/*
https://www.chessprogramming.org/Pawn_Pattern_and_Properties
https://www.chessprogramming.org/Pawn_Pushes_(Bitboards)
https://www.chessprogramming.org/Pawn_Attacks_(Bitboards)
*/

import (
	"math/bits"

	"github.com/geoffroyp2/chessgo/engine/bitboard"
	"github.com/geoffroyp2/chessgo/engine/constants"
	"github.com/geoffroyp2/chessgo/engine/move"
	"github.com/geoffroyp2/chessgo/engine/position"
)

var wPawnCapturesLookup = [64]uint64{
	0x2<<(8*1), 0x5<<(8*1), 0xA<<(8*1), 0x14<<(8*1), 0x28<<(8*1), 0x50<<(8*1), 0xA0<<(8*1), 0x40<<(8*1), // Rank 1
	0x2<<(8*2), 0x5<<(8*2), 0xA<<(8*2), 0x14<<(8*2), 0x28<<(8*2), 0x50<<(8*2), 0xA0<<(8*2), 0x40<<(8*2), // Rank 2
	0x2<<(8*3), 0x5<<(8*3), 0xA<<(8*3), 0x14<<(8*3), 0x28<<(8*3), 0x50<<(8*3), 0xA0<<(8*3), 0x40<<(8*3), // Rank 3
	0x2<<(8*4), 0x5<<(8*4), 0xA<<(8*4), 0x14<<(8*4), 0x28<<(8*4), 0x50<<(8*4), 0xA0<<(8*4), 0x40<<(8*4), // Rank 4
	0x2<<(8*5), 0x5<<(8*5), 0xA<<(8*5), 0x14<<(8*5), 0x28<<(8*5), 0x50<<(8*5), 0xA0<<(8*5), 0x40<<(8*5), // Rank 5
	0x2<<(8*6), 0x5<<(8*6), 0xA<<(8*6), 0x14<<(8*6), 0x28<<(8*6), 0x50<<(8*6), 0xA0<<(8*6), 0x40<<(8*6), // Rank 6
	0x2<<(8*7), 0x5<<(8*7), 0xA<<(8*7), 0x14<<(8*7), 0x28<<(8*7), 0x50<<(8*7), 0xA0<<(8*7), 0x40<<(8*7), // Rank 7
	0,          0,          0,          0,           0,           0,           0,           0,           // Rank 8
}
var bPawnCapturesLookup = [64]uint64{
	0,          0,          0,          0,           0,           0,           0,           0,           // Rank 1
	0x2<<(8*0), 0x5<<(8*0), 0xA<<(8*0), 0x14<<(8*0), 0x28<<(8*0), 0x50<<(8*0), 0xA0<<(8*0), 0x40<<(8*0), // Rank 2
	0x2<<(8*1), 0x5<<(8*1), 0xA<<(8*1), 0x14<<(8*1), 0x28<<(8*1), 0x50<<(8*1), 0xA0<<(8*1), 0x40<<(8*1), // Rank 3
	0x2<<(8*2), 0x5<<(8*2), 0xA<<(8*2), 0x14<<(8*2), 0x28<<(8*2), 0x50<<(8*2), 0xA0<<(8*2), 0x40<<(8*2), // Rank 4
	0x2<<(8*3), 0x5<<(8*3), 0xA<<(8*3), 0x14<<(8*3), 0x28<<(8*3), 0x50<<(8*3), 0xA0<<(8*3), 0x40<<(8*3), // Rank 5
	0x2<<(8*4), 0x5<<(8*4), 0xA<<(8*4), 0x14<<(8*4), 0x28<<(8*4), 0x50<<(8*4), 0xA0<<(8*4), 0x40<<(8*4), // Rank 6
	0x2<<(8*5), 0x5<<(8*5), 0xA<<(8*5), 0x14<<(8*5), 0x28<<(8*5), 0x50<<(8*5), 0xA0<<(8*5), 0x40<<(8*5), // Rank 7
	0x2<<(8*6), 0x5<<(8*6), 0xA<<(8*6), 0x14<<(8*6), 0x28<<(8*6), 0x50<<(8*6), 0xA0<<(8*6), 0x40<<(8*6), // Rank 8
}

func wPawnMoves(pos *position.Position, moveArray *constants.MoveArray, moveAmount int) int {

	// Quiet moves
	pushablePawns := bitboard.ShiftS(pos.Empty) & pos.Pieces[constants.WHITEPAWN]
	
	normalMoves := pushablePawns & constants.NotRank7
	for normalMoves != 0 {
		pawnIdx := uint32(bits.TrailingZeros64(normalMoves))
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx + 8, move.MOVE, constants.WHITEPAWN, constants.NULLPIECE)
		moveAmount++
		normalMoves ^= 1<<pawnIdx
	}

	dblMoves := pushablePawns & bitboard.ShiftS(bitboard.ShiftS(pos.Empty & constants.Rank4)) // Empty rank 4 shifted 2 times
	for dblMoves != 0 {
		pawnIdx := uint32(bits.TrailingZeros64(dblMoves))
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx + 16, move.PMOVE2, constants.WHITEPAWN, constants.NULLPIECE)
		moveAmount++
		dblMoves ^= 1<<pawnIdx
	}

	promotions := pushablePawns & constants.Rank7
	for promotions != 0 {
		pawnIdx := uint32(bits.TrailingZeros64(promotions))
		(*moveArray)[moveAmount]     = move.CreateMove(pawnIdx, pawnIdx + 8, move.QPROM, constants.WHITEPAWN, constants.NULLPIECE)
		(*moveArray)[moveAmount + 1] = move.CreateMove(pawnIdx, pawnIdx + 8, move.RPROM, constants.WHITEPAWN, constants.NULLPIECE)
		(*moveArray)[moveAmount + 2] = move.CreateMove(pawnIdx, pawnIdx + 8, move.BPROM, constants.WHITEPAWN, constants.NULLPIECE)
		(*moveArray)[moveAmount + 3] = move.CreateMove(pawnIdx, pawnIdx + 8, move.NPROM, constants.WHITEPAWN, constants.NULLPIECE)
		moveAmount += 4
		promotions ^= 1<<pawnIdx
	}

	// Captures
	// All opponents pieces + En-passant square (==> Shift by 2 more then back to discard set bit if EPSquare is not set (= 63))
	captureTargets := pos.AllBlack | uint64((1 << (pos.GetEPSquare() + 2)) >> 2)
	
	normalPawns := pos.Pieces[constants.WHITEPAWN] & constants.NotRank7
	for normalPawns != 0 {
		pawnIdx := uint32(bits.TrailingZeros64(normalPawns))
		captures := wPawnCapturesLookup[pawnIdx] & captureTargets

		for captures != 0 {
			destIdx := uint32(bits.TrailingZeros64(captures))
			// initialized to BLACKPAWN & EPCAPTURE for en-passant (because the capture square is empty)
			capturedPiece := constants.BLACKPAWN 
			moveType := move.EPCAPTURE
			for pi := constants.BLACKKING; pi <= constants.BLACKPAWN; pi++ {
				if pos.Pieces[pi] & (1<<destIdx) != 0 {
					capturedPiece = pi
					moveType = move.CAPTURE
					break
				}
			}
			(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, destIdx, moveType, constants.WHITEPAWN, capturedPiece)
			moveAmount++
			captures ^= 1<<destIdx
		}

		normalPawns ^= 1<<pawnIdx
	}

	promPawns := pos.Pieces[constants.WHITEPAWN] & constants.Rank7
	for promPawns != 0 {
		pawnIdx := uint32(bits.TrailingZeros64(promPawns))
		captures := wPawnCapturesLookup[pawnIdx] & captureTargets
		
		for captures != 0 {
			destIdx := uint32(bits.TrailingZeros64(captures))
			capturedPiece := getBCapturedPiece(pos, destIdx)

			(*moveArray)[moveAmount]     = move.CreateMove(pawnIdx, destIdx, move.QPROMCAPTURE, constants.WHITEPAWN, capturedPiece)
			(*moveArray)[moveAmount + 1] = move.CreateMove(pawnIdx, destIdx, move.RPROMCAPTURE, constants.WHITEPAWN, capturedPiece)
			(*moveArray)[moveAmount + 2] = move.CreateMove(pawnIdx, destIdx, move.BPROMCAPTURE, constants.WHITEPAWN, capturedPiece)
			(*moveArray)[moveAmount + 3] = move.CreateMove(pawnIdx, destIdx, move.NPROMCAPTURE, constants.WHITEPAWN, capturedPiece)
			moveAmount += 4
			captures ^= 1<<destIdx
		}

		promPawns ^= 1<<pawnIdx
	}

	return moveAmount
}

func bPawnMoves(pos *position.Position, moveArray *constants.MoveArray, moveAmount int) int {
	// Quiet moves
	pushablePawns := bitboard.ShiftN(pos.Empty) & pos.Pieces[constants.BLACKPAWN]
	
	normalMoves := pushablePawns & constants.NotRank2
	for normalMoves != 0 {
		pawnIdx := uint32(bits.TrailingZeros64(normalMoves))
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx - 8, move.MOVE, constants.BLACKPAWN, constants.NULLPIECE)
		moveAmount++
		normalMoves ^= 1<<pawnIdx
	}

	dblMoves := pushablePawns & bitboard.ShiftN(bitboard.ShiftN(pos.Empty & constants.Rank5))
	for dblMoves != 0 {
		pawnIdx := uint32(bits.TrailingZeros64(dblMoves))
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx - 16, move.PMOVE2, constants.BLACKPAWN, constants.NULLPIECE)
		moveAmount++
		dblMoves ^= 1<<pawnIdx
	}

	promotions := pushablePawns & constants.Rank2
	for promotions != 0 {
		pawnIdx := uint32(bits.TrailingZeros64(promotions))
		(*moveArray)[moveAmount]     = move.CreateMove(pawnIdx, pawnIdx - 8, move.QPROM, constants.BLACKPAWN, constants.NULLPIECE)
		(*moveArray)[moveAmount + 1] = move.CreateMove(pawnIdx, pawnIdx - 8, move.RPROM, constants.BLACKPAWN, constants.NULLPIECE)
		(*moveArray)[moveAmount + 2] = move.CreateMove(pawnIdx, pawnIdx - 8, move.BPROM, constants.BLACKPAWN, constants.NULLPIECE)
		(*moveArray)[moveAmount + 3] = move.CreateMove(pawnIdx, pawnIdx - 8, move.NPROM, constants.BLACKPAWN, constants.NULLPIECE)
		moveAmount += 4
		promotions ^= 1<<pawnIdx
	}

	// Captures
	captureTargets := pos.AllWhite | uint64((1 << (pos.GetEPSquare() + 2)) >> 2)
	
	normalPawns := pos.Pieces[constants.BLACKPAWN] & constants.NotRank2
	for normalPawns != 0 {
		pawnIdx := uint32(bits.TrailingZeros64(normalPawns))
		captures := bPawnCapturesLookup[pawnIdx] & captureTargets

		for captures != 0 {
			destIdx := uint32(bits.TrailingZeros64(captures))
			capturedPiece := constants.WHITEPAWN
			moveType := move.EPCAPTURE
			for pi := constants.WHITEKING; pi <= constants.WHITEPAWN; pi++ {
				if pos.Pieces[pi] & (1<<destIdx) != 0 {
					capturedPiece = pi
					moveType = move.CAPTURE
					break
				}
			}
			(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, destIdx, moveType, constants.BLACKPAWN, capturedPiece)
			moveAmount++
			captures ^= 1<<destIdx
		}

		normalPawns ^= 1<<pawnIdx
	}

	promPawns := pos.Pieces[constants.BLACKPAWN] & constants.Rank2
	for promPawns != 0 {
		pawnIdx := uint32(bits.TrailingZeros64(promPawns))
		captures := bPawnCapturesLookup[pawnIdx] & captureTargets
		
		for captures != 0 {
			destIdx := uint32(bits.TrailingZeros64(captures))
			capturedPiece := getWCapturedPiece(pos, destIdx)

			(*moveArray)[moveAmount]     = move.CreateMove(pawnIdx, destIdx, move.QPROMCAPTURE, constants.BLACKPAWN, capturedPiece)
			(*moveArray)[moveAmount + 1] = move.CreateMove(pawnIdx, destIdx, move.RPROMCAPTURE, constants.BLACKPAWN, capturedPiece)
			(*moveArray)[moveAmount + 2] = move.CreateMove(pawnIdx, destIdx, move.BPROMCAPTURE, constants.BLACKPAWN, capturedPiece)
			(*moveArray)[moveAmount + 3] = move.CreateMove(pawnIdx, destIdx, move.NPROMCAPTURE, constants.BLACKPAWN, capturedPiece)
			moveAmount += 4
			captures ^= 1<<destIdx
		}

		promPawns ^= 1<<pawnIdx
	}

	return moveAmount
}

// Generate a simple BitBoard of squares that can be attacked by pawns
func wPawnAttacks(pos *position.Position) uint64 {
	var attacks uint64 = 0
	pawns := pos.Pieces[constants.WHITEPAWN]
	for pawns != 0 {
		pawnIdx := bits.TrailingZeros64(pawns)
		attacks |= wPawnCapturesLookup[pawnIdx]
		pawns ^= 1<<pawnIdx
	}
	return attacks
}

func bPawnAttacks(pos *position.Position) uint64 {
	var attacks uint64 = 0
	pawns := pos.Pieces[constants.BLACKPAWN]
	for pawns != 0 {
		pawnIdx := bits.TrailingZeros64(pawns)
		attacks |= bPawnCapturesLookup[pawnIdx]
		pawns ^= 1<<pawnIdx
	}
	return attacks
}
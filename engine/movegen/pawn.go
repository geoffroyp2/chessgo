package movegen

/*
	TODO: change captures to use precomputed patterns
	https://www.chessprogramming.org/Pawn_Attacks_(Bitboards)
*/

import (
	"math/bits"

	"github.com/geoffroyp2/chessgo/engine/bitboard"
	"github.com/geoffroyp2/chessgo/engine/constants"
	"github.com/geoffroyp2/chessgo/engine/move"
	"github.com/geoffroyp2/chessgo/engine/position"
)

func wPawnMoves(pos *position.Position, moveArray *constants.MoveArray, moveAmount int) int {
	moveAmount = wMoves(pos, moveArray, moveAmount)
	moveAmount = wCaptures(pos, moveArray, moveAmount)
	return moveAmount
}

func bPawnMoves(pos *position.Position, moveArray *constants.MoveArray, moveAmount int) int {
	moveAmount = bMoves(pos, moveArray, moveAmount)		
	moveAmount = bCaptures(pos, moveArray, moveAmount)
	return moveAmount
}

func wMoves(pos *position.Position, moveArray *constants.MoveArray, moveAmount int) int {
	// Intersection of empty squares shifted down once and set of pawns
	pushablePawns := bitboard.ShiftS(pos.Empty) & pos.Pieces[constants.WHITEPAWN]
	
	dblMoves := pushablePawns & bitboard.ShiftS(bitboard.ShiftS(pos.Empty & constants.Rank4)) // Empty rank 4 shifted 2 times
	promotions := pushablePawns & constants.Rank7
	normalMoves := pushablePawns & ^constants.Rank7

	// Iterate over set bits and generate the associated move
	for ; normalMoves != 0 ; {
		pawnIdx := uint32(bits.TrailingZeros64(normalMoves))
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx + 8, move.MOVE, constants.WHITEPAWN, constants.NULLPIECE)
		moveAmount++
		normalMoves ^= 1 << pawnIdx
	}

	// Double Push
	for ; dblMoves != 0 ; {
		pawnIdx := uint32(bits.TrailingZeros64(dblMoves))
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx + 16, move.PMOVE2, constants.WHITEPAWN, constants.NULLPIECE)
		moveAmount++
		dblMoves ^= 1 << pawnIdx
	}
	
	// Promotions
	for ; promotions != 0 ; {
		pawnIdx := uint32(bits.TrailingZeros64(promotions))
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx + 8, move.QPROM, constants.WHITEPAWN, constants.NULLPIECE)
		moveAmount++
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx + 8, move.RPROM, constants.WHITEPAWN, constants.NULLPIECE)
		moveAmount++
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx + 8, move.BPROM, constants.WHITEPAWN, constants.NULLPIECE)
		moveAmount++
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx + 8, move.NPROM, constants.WHITEPAWN, constants.NULLPIECE)
		moveAmount++
		promotions ^= 1 << pawnIdx
	}

	return moveAmount
}

func wCaptures(pos *position.Position, moveArray *constants.MoveArray, moveAmount int) int {

	// Shift all black pieces in the reverse direction of the pawn capture & intersect with set of pawns
	captureR := bitboard.ShiftSW(pos.AllBlack) & pos.Pieces[constants.WHITEPAWN]
	captureL := bitboard.ShiftSE(pos.AllBlack) & pos.Pieces[constants.WHITEPAWN]

	// Separate normal captures and promotion-captures
	normalR := captureR & ^constants.Rank7
	promotionR := captureR & constants.Rank7
	normalL := captureL & ^constants.Rank7
	promotionL := captureL & constants.Rank7

	// Find en-passant captures: get EPSquare, convert it to a BitBoard then intersect with pawns that can capture it
	var EPSquareBB uint64 = 0
	EPSquare := pos.GetEPSquare()
	if EPSquare != constants.EPSQUARE_NONE {
		EPSquareBB = 1 << EPSquare
	}
	EPCaptureR := bitboard.ShiftSW(EPSquareBB) & pos.Pieces[constants.WHITEPAWN]
	EPCaptureL := bitboard.ShiftSE(EPSquareBB) & pos.Pieces[constants.WHITEPAWN]

	// Normal captures
	for ; normalR != 0 ; {
		pawnIdx := uint32(bits.TrailingZeros64(normalR))

		// Find the captured piece's type
		capturedPiece := constants.NULLPIECE
		for pi := constants.BLACKKING; pi <= constants.BLACKPAWN; pi++ {
			if bitboard.ShiftSW(pos.Pieces[pi]) & (1 << pawnIdx) != 0 {
				capturedPiece = pi
			}
		}

		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx + 9, move.CAPTURE, constants.WHITEPAWN, capturedPiece)
		moveAmount++
		normalR ^= 1 << pawnIdx
	}

	for ; normalL != 0 ; {
		pawnIdx := uint32(bits.TrailingZeros64(normalL))

		capturedPiece := constants.NULLPIECE
		for pi := constants.BLACKKING; pi <= constants.BLACKPAWN; pi++ {
			if bitboard.ShiftSE(pos.Pieces[pi]) & (1 << pawnIdx) != 0 {
				capturedPiece = pi
			}
		}

		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx + 7, move.CAPTURE, constants.WHITEPAWN, capturedPiece)
		moveAmount++
		normalL ^= 1 << pawnIdx
	}

	// Capture en-passant
	for ; EPCaptureR != 0 ; {
		pawnIdx := uint32(bits.TrailingZeros64(EPCaptureR))
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx + 9, move.EPCAPTURE, constants.WHITEPAWN, constants.BLACKPAWN)
		moveAmount++
		EPCaptureR ^= 1 << pawnIdx
	}

	for ; EPCaptureL != 0 ; {
		pawnIdx := uint32(bits.TrailingZeros64(EPCaptureL))
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx + 7, move.EPCAPTURE, constants.WHITEPAWN, constants.BLACKPAWN)
		moveAmount++
		EPCaptureL ^= 1 << pawnIdx
	}

	// Promotion - capture
	for ; promotionR != 0 ; {
		pawnIdx := uint32(bits.TrailingZeros64(promotionR))

		capturedPiece := constants.NULLPIECE
		for pi := constants.BLACKKING; pi <= constants.BLACKPAWN; pi++ {
			if bitboard.ShiftSW(pos.Pieces[pi]) & (1 << pawnIdx) != 0 {
				capturedPiece = pi
			}
		}

		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx + 9, move.QPROMCAPTURE, constants.WHITEPAWN, capturedPiece)
		moveAmount++
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx + 9, move.RPROMCAPTURE, constants.WHITEPAWN, capturedPiece)
		moveAmount++
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx + 9, move.BPROMCAPTURE, constants.WHITEPAWN, capturedPiece)
		moveAmount++
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx + 9, move.NPROMCAPTURE, constants.WHITEPAWN, capturedPiece)
		moveAmount++
		promotionR ^= 1 << pawnIdx
	}

	for ; promotionL != 0 ; {
		pawnIdx := uint32(bits.TrailingZeros64(promotionL))

		capturedPiece := constants.NULLPIECE
		for pi := constants.BLACKKING; pi <= constants.BLACKPAWN; pi++ {
			if bitboard.ShiftSE(pos.Pieces[pi]) & (1 << pawnIdx) != 0 {
				capturedPiece = pi
			}
		}

		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx + 7, move.QPROMCAPTURE, constants.WHITEPAWN, capturedPiece)
		moveAmount++
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx + 7, move.RPROMCAPTURE, constants.WHITEPAWN, capturedPiece)
		moveAmount++
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx + 7, move.BPROMCAPTURE, constants.WHITEPAWN, capturedPiece)
		moveAmount++
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx + 7, move.NPROMCAPTURE, constants.WHITEPAWN, capturedPiece)
		moveAmount++
		promotionL ^= 1 << pawnIdx
	}

	return moveAmount
}

func bMoves(pos *position.Position, moveArray *constants.MoveArray, moveAmount int) int {
	// Intersection of empty squares shifted up once and set of pawns
	pushablePawns := bitboard.ShiftN(pos.Empty) & pos.Pieces[constants.BLACKPAWN]
	dblPushable := pushablePawns & bitboard.ShiftN(bitboard.ShiftN(pos.Empty & constants.Rank5)) // Empty rank 5 shifted 2 times
	promotions := pushablePawns & constants.Rank2

	// Iterate over set bits and generate the associated move
	for ; pushablePawns != 0 ; {
		pawnIdx := uint32(bits.TrailingZeros64(pushablePawns))
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx - 8, move.MOVE, constants.BLACKPAWN, constants.NULLPIECE)
		moveAmount++
		pushablePawns ^= 1 << pawnIdx
	}

	// Double Push
	for ; dblPushable != 0 ; {
		pawnIdx := uint32(bits.TrailingZeros64(dblPushable))
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx - 16, move.PMOVE2, constants.BLACKPAWN, constants.NULLPIECE)
		moveAmount++
		dblPushable ^= 1 << pawnIdx
	}
	
	// Promotions
	for ; promotions != 0 ; {
		pawnIdx := uint32(bits.TrailingZeros64(promotions))
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx - 8, move.QPROM, constants.BLACKPAWN, constants.NULLPIECE)
		moveAmount++
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx - 8, move.RPROM, constants.BLACKPAWN, constants.NULLPIECE)
		moveAmount++
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx - 8, move.BPROM, constants.BLACKPAWN, constants.NULLPIECE)
		moveAmount++
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx - 8, move.NPROM, constants.BLACKPAWN, constants.NULLPIECE)
		moveAmount++
		promotions ^= 1 << pawnIdx
	}

	return moveAmount
}

func bCaptures(pos *position.Position, moveArray *constants.MoveArray, moveAmount int) int {

	// Shift all black pieces in the reverse direction of the pawn capture & intersect with set of pawns
	captureR := bitboard.ShiftNW(pos.AllWhite) & pos.Pieces[constants.BLACKPAWN]
	captureL := bitboard.ShiftNE(pos.AllWhite) & pos.Pieces[constants.BLACKPAWN]

	// Separate normal captures and promotion-captures
	normalR := captureR & ^constants.Rank2
	promotionsR := captureR & constants.Rank2
	normalL := captureL & ^constants.Rank2
	promotionsL := captureL & constants.Rank2

	// Find en-passant captures: get EPSquare, convert it to a BitBoard then intersect with pawns that can capture it
	var EPSquareBB uint64 = 0
	EPSquare := pos.GetEPSquare()
	if EPSquare != constants.EPSQUARE_NONE {
		EPSquareBB = 1 << EPSquare
	}
	EPCaptureR := bitboard.ShiftNW(EPSquareBB) & pos.Pieces[constants.BLACKPAWN]
	EPCaptureL := bitboard.ShiftNE(EPSquareBB) & pos.Pieces[constants.BLACKPAWN]

	// Normal captures
	for ; normalR != 0 ; {
		pawnIdx := uint32(bits.TrailingZeros64(normalR))

		// Find the captured piece's type
		capturedPiece := constants.NULLPIECE
		for pi := constants.WHITEKING; pi <= constants.WHITEPAWN; pi++ {
			if bitboard.ShiftNW(pos.Pieces[pi]) & (1 << pawnIdx) != 0 {
				capturedPiece = pi
			}
		}

		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx - 7, move.CAPTURE, constants.BLACKPAWN, capturedPiece)
		moveAmount++
		normalR ^= 1 << pawnIdx
	}

	for ; normalL != 0 ; {
		pawnIdx := uint32(bits.TrailingZeros64(normalL))

		capturedPiece := constants.NULLPIECE
		for pi := constants.WHITEKING; pi <= constants.WHITEPAWN; pi++ {
			if bitboard.ShiftNE(pos.Pieces[pi]) & (1 << pawnIdx) != 0 {
				capturedPiece = pi
			}
		}

		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx - 9, move.CAPTURE, constants.BLACKPAWN, capturedPiece)
		moveAmount++
		normalL ^= 1 << pawnIdx
	}
	
	// Capture en-passant
	for ; EPCaptureR != 0 ; {
		pawnIdx := uint32(bits.TrailingZeros64(EPCaptureR))
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx - 7, move.EPCAPTURE, constants.BLACKPAWN, constants.WHITEPAWN)
		moveAmount++
		EPCaptureR ^= 1 << pawnIdx
	}

	for ; EPCaptureL != 0 ; {
		pawnIdx := uint32(bits.TrailingZeros64(EPCaptureL))
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx - 9, move.EPCAPTURE, constants.BLACKPAWN, constants.WHITEPAWN)
		moveAmount++
		EPCaptureL ^= 1 << pawnIdx
	}

	// Promotion - capture
	for ; promotionsR != 0 ; {
		pawnIdx := uint32(bits.TrailingZeros64(promotionsR))

		capturedPiece := constants.NULLPIECE
		for pi := constants.WHITEKING; pi <= constants.WHITEPAWN; pi++ {
			if bitboard.ShiftNW(pos.Pieces[pi]) & (1 << pawnIdx) != 0 {
				capturedPiece = pi
			}
		}

		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx - 7, move.QPROMCAPTURE, constants.BLACKPAWN, capturedPiece)
		moveAmount++
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx - 7, move.RPROMCAPTURE, constants.BLACKPAWN, capturedPiece)
		moveAmount++
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx - 7, move.BPROMCAPTURE, constants.BLACKPAWN, capturedPiece)
		moveAmount++
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx - 7, move.NPROMCAPTURE, constants.BLACKPAWN, capturedPiece)
		moveAmount++
		promotionsR ^= 1 << pawnIdx
	}

	for ; promotionsL != 0 ; {
		pawnIdx := uint32(bits.TrailingZeros64(promotionsL))

		capturedPiece := constants.NULLPIECE
		for pi := constants.WHITEKING; pi <= constants.WHITEPAWN; pi++ {
			if bitboard.ShiftNE(pos.Pieces[pi]) & (1 << pawnIdx) != 0 {
				capturedPiece = pi
			}
		}

		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx - 9, move.QPROMCAPTURE, constants.BLACKPAWN, capturedPiece)
		moveAmount++
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx - 9, move.RPROMCAPTURE, constants.BLACKPAWN, capturedPiece)
		moveAmount++
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx - 9, move.BPROMCAPTURE, constants.BLACKPAWN, capturedPiece)
		moveAmount++
		(*moveArray)[moveAmount] = move.CreateMove(pawnIdx, pawnIdx - 9, move.NPROMCAPTURE, constants.BLACKPAWN, capturedPiece)
		moveAmount++
		promotionsL ^= 1 << pawnIdx
	}

	return moveAmount
}
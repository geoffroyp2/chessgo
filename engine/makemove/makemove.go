package makemove

import (
	"github.com/geoffroyp2/chessgo/engine/constants"
	"github.com/geoffroyp2/chessgo/engine/move"
	"github.com/geoffroyp2/chessgo/engine/position"
)

func MakeMove(pos *position.Position, m uint32) {

	from := move.From(m)
	to := move.To(m)
	mtype := move.Type(m)
	piece := move.MainPiece(m)
	extra := move.ExtraPiece(m)
	turn := pos.GetPlayerTurn()

	pos.SetEPSquare(constants.EPSQUARE_NONE)

	switch mtype {
	case move.MOVE:
		pos.Pieces[piece] ^= (1<<from) | (1<<to) // Remove piece at from square and add it back at to square
		break
	case move.PMOVE2: 
		pos.Pieces[piece] ^= (1<<from) | (1<<to)
		if turn == constants.PLAYERWHITE {
			pos.SetEPSquare(uint16(from + 8)) // Set EPSquare
		} else {
			pos.SetEPSquare(uint16(from - 8))
		}
		break
	case move.KCASTLE:
		pos.Pieces[piece] ^= (1<<from) | (1<<to) // King
		pos.Pieces[extra] ^= (1<<(to+1)) | (1<<(to-1)) // Rook
		pos.RemoveCastleRights(turn)
		break
	case move.QCASTLE:
		pos.Pieces[piece] ^= (1<<from) | (1<<to) // King
		pos.Pieces[extra] ^= (1<<(to-2)) | (1<<(to+1)) // Rook
		pos.RemoveCastleRights(turn)
		break
	case move.CAPTURE:
		pos.Pieces[piece] ^= (1<<from) | (1<<to) // Move piece
		pos.Pieces[extra] ^= 1<<to // Remove captured piece
		break
	case move.EPCAPTURE:
		pos.Pieces[piece] ^= (1<<from) | (1<<to) // Move piece
		if turn == constants.PLAYERWHITE {
			pos.Pieces[extra] ^= 1<<(to-8)
		} else {
			pos.Pieces[extra] ^= 1<<(to+8)
		}
		break
	case move.QPROM: fallthrough
	case move.RPROM: fallthrough
	case move.BPROM: fallthrough
	case move.NPROM:
		pos.Pieces[piece] ^= 1<<from // Remove piece
		pos.Pieces[move.GetPromotionTarget(m) + 6*uint32(turn)] ^= 1<<to // Add promoted piece
		break
	case move.QPROMCAPTURE: fallthrough
	case move.RPROMCAPTURE: fallthrough
	case move.BPROMCAPTURE: fallthrough
	case move.NPROMCAPTURE:
		pos.Pieces[piece] ^= 1<<from // Remove piece
		pos.Pieces[extra] ^= 1<<to // Remove captured piece
		pos.Pieces[move.GetPromotionTarget(m) + 6*uint32(turn)] ^= 1<<to // Add promoted piece
		break
	}

	// Half-move clock
	if piece == constants.WHITEPAWN || piece == constants.BLACKPAWN || move.IsCapture(m) {
		pos.HMClock = 0
	} else {
		pos.HMClock++
	}

	pos.TogglePLayerTurn()
	pos.ComputeUnions() // TODO: only compute new move
}
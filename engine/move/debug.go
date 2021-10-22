package move

import (
	"fmt"

	"github.com/geoffroyp2/chessgo/engine/constants"
	"github.com/geoffroyp2/chessgo/engine/stringutils"
)

func PrintMove(move uint32) {

	mFrom := stringutils.IdxToCoordString(uint16(From(move)))
	mTo := stringutils.IdxToCoordString(uint16(To(move)))
	typeString := getMoveTypeString(Type(move))
	mainPiece := getPieceString(MainPiece(move))
	extraPiece := getPieceString(ExtraPiece(move))

	fmt.Printf("%s %s%s %s %s\n", mainPiece, *mFrom, *mTo, typeString, extraPiece)
}

func getMoveTypeString(t uint32) string {
	switch t {
		case MOVE:         return "Move"
		case PMOVE2:       return "Double move"
		case KCASTLE:      return "Castle King-side"
		case QCASTLE:      return "Castle Queen-side"
		case CAPTURE:      return "Capture"
		case EPCAPTURE:    return "Capture en-passant"
		case QPROM:        return "Promotion Queen"
		case RPROM:        return "Promotion Rook"
		case BPROM:        return "Promotion Bishop"
		case NPROM:        return "Promotion Knight"
		case QPROMCAPTURE: return "Promotion Queen & Capture"
		case RPROMCAPTURE: return "Promotion Rook & Capture"
		case BPROMCAPTURE: return "Promotion Bishop & Capture"
		case NPROMCAPTURE: return "Promotion Knight & Capture"
	}
	return ""
}

func getPieceString(p uint32) string {
	switch p {
		case constants.WHITEKING:   return "White King"
		case constants.WHITEQUEEN:  return "White Queen"
		case constants.WHITEROOK:   return "White Rook"
		case constants.WHITEBISHOP: return "White Bishop"
		case constants.WHITEKNIGHT: return "White Knight"
		case constants.WHITEPAWN:   return "White Pawn"
		case constants.BLACKKING:   return "Black King"
		case constants.BLACKQUEEN:  return "Black Queen"
		case constants.BLACKROOK:   return "Black Rook"
		case constants.BLACKBISHOP: return "Black Bishop"
		case constants.BLACKKNIGHT: return "Black Knight"
		case constants.BLACKPAWN:   return "Black Pawn"
	}
	return ""
}
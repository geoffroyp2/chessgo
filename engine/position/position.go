package position

import "github.com/geoffroyp2/chessgo/engine/constants"


type Position struct {
	Pieces [12]uint64 // All pieces as bitboards + unions of all black and all white pieces
	AllWhite uint64
	AllBlack uint64
	Occupied uint64
	Empty uint64
	Flags uint16 			// Flags for player turn, castle rights and en-passant square
	HMClock uint16 		// Half Move Clock (50 max is a draw)
	MoveNumber uint32 
}

func NewPosition() *Position {
	return &Position{}
}

func (pos *Position) GetPieceBB(pType constants.PieceType) uint64 {
	return pos.Pieces[pType]
}

func (pos *Position) SetPieceBB(pType constants.PieceType, newValue uint64) {
	pos.Pieces[pType] = newValue
}


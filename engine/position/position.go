package position

type Position struct {
	Pieces   [12]uint64 // All pieces as bitboards + unions of all black and all white pieces
	AllWhite uint64
	AllBlack uint64
	Occupied uint64
	Empty    uint64
	Flags    uint16 // Flags for player turn, castle rights and en-passant square
	HMClock  uint16 // Half Move Clock (50 max is a draw)
}

func NewPosition() *Position {
	return &Position{}
}

func CopyPosition(pos *Position) *Position {
	newPos := NewPosition()
	for i := 0; i < 12; i++ {
		newPos.Pieces[i] = pos.Pieces[i]
	}
	newPos.AllBlack = pos.AllBlack
	newPos.AllWhite = pos.AllWhite
	newPos.Occupied = pos.Occupied
	newPos.Empty = pos.Empty
	newPos.Flags = pos.Flags
	newPos.HMClock = pos.HMClock
	return newPos
}

func (pos *Position) GetPieceBB(pType uint32) uint64 {
	return pos.Pieces[pType]
}

func (pos *Position) SetPieceBB(pType uint32, newValue uint64) {
	pos.Pieces[pType] = newValue
}

func (pos *Position) ComputeUnions() {
	pos.AllWhite = 0
	for idx := 0; idx < 6; idx++ {
		pos.AllWhite |= pos.Pieces[idx]
	}
	pos.AllBlack = 0
	for idx := 6; idx < 12; idx++ {
		pos.AllBlack |= pos.Pieces[idx]
	}
	pos.Occupied = pos.AllWhite | pos.AllBlack
	pos.Empty = ^pos.Occupied
}
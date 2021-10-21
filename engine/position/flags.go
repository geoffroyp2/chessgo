package position

const (
	pt_o     uint16 = 0           // Player turn offset in flags
	pt_mask  uint16 = 0x1 << pt_o // 1 bit toggled for player turn
	cr_o     uint16 = 1
	cr_mask  uint16 = 0xF << cr_o // 4 bits for castle rights KQkq
	eps_o    uint16 = 5
	eps_mask uint16 = 0x3F << eps_o // 6 bits for en-passant square 0-63
)

// Player turn
func (pos *Position) GetPlayerTurn() uint16 {
	return pos.Flags & pt_mask
}

func (pos *Position) SetPlayerTurn(newVal uint16) {
	pos.Flags = (pos.Flags & ^pt_mask) | newVal
}

func (pos *Position) TogglePLayerTurn() {
	pos.Flags ^= 1
}

// Castle Rights
func (pos *Position) GetCastleRights() uint16 {
	return (pos.Flags & cr_mask) >> cr_o
}

func (pos *Position) GetCastleRight(idx uint16 /*KQkq*/) bool {
	return pos.Flags&(1<<(cr_o+4-idx)) != 0
}

func (pos *Position) SetCastleRights(newVal uint16) {
	pos.Flags = (pos.Flags & ^cr_mask) | (newVal << cr_o)
}

func (pos *Position) ToggleCastleRigth(idx uint16 /*KQkq*/) {
	pos.Flags ^= 1 << (cr_o + 3 - idx)
}

// En-passant square
func (pos *Position) GetEPSquare() uint16 {
	return (pos.Flags & eps_mask) >> eps_o
}

func (pos *Position) SetEPSquare(newVal uint16) {
	pos.Flags = (pos.Flags & ^eps_mask) | (newVal << eps_o)
}
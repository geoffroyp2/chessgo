package constants

type MoveArray [250]uint32 // 218 is the maximum moves amount for a given position, 250 to be safe

const (
	WHITEKING   uint32 = iota
	WHITEQUEEN  uint32 = iota
	WHITEROOK   uint32 = iota
	WHITEBISHOP uint32 = iota
	WHITEKNIGHT uint32 = iota
	WHITEPAWN   uint32 = iota
	BLACKKING   uint32 = iota
	BLACKQUEEN  uint32 = iota
	BLACKROOK   uint32 = iota
	BLACKBISHOP uint32 = iota
	BLACKKNIGHT uint32 = iota
	BLACKPAWN   uint32 = iota
	NULLPIECE   uint32 = 0xF
)

const EPSQUARE_NONE uint16 = 0x3F // Represents the value if there is no en-passant square in a position

const NotAFile uint64 = 0xfefefefefefefefe
const NotHFile uint64 = 0x7f7f7f7f7f7f7f7f
const Rank1 uint64 = 0x00000000000000ff
const Rank2 uint64 = 0x000000000000ff00
const Rank3 uint64 = 0x0000000000ff0000
const Rank4 uint64 = 0x00000000ff000000
const Rank5 uint64 = 0x000000ff00000000
const Rank6 uint64 = 0x0000ff0000000000
const Rank7 uint64 = 0x00ff000000000000
const Rank8 uint64 = 0xff00000000000000

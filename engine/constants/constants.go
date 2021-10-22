package constants

type MoveArray [250]uint32 // 218 is the maximum moves amount for a given position, 250 to be safe

// Piece types
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

// Directions
const (
	DIRN  = iota
	DIRS  = iota
	DIRE  = iota
	DIRW  = iota
	DIRNE = iota
	DIRNW = iota
	DIRSE = iota
	DIRSW = iota
)
const DIRSTRAIGHT = 0
const DIRDIAGONAL = 4

// Others
const EPSQUARE_NONE uint16 = 0x3F // Represents the value if there is no en-passant square in a position

// Ranks and files masks
const FileA uint64 = 0x101010101010101
const FileB uint64 = 0x202020202020202
const FileC uint64 = 0x404040404040404
const FileD uint64 = 0x808080808080808
const FileE uint64 = 0x1010101010101010
const FileF uint64 = 0x2020202020202020
const FileG uint64 = 0x4040404040404040
const FileH uint64 = 0x8080808080808080
const Rank1 uint64 = 0x00000000000000ff
const Rank2 uint64 = 0x000000000000ff00
const Rank3 uint64 = 0x0000000000ff0000
const Rank4 uint64 = 0x00000000ff000000
const Rank5 uint64 = 0x000000ff00000000
const Rank6 uint64 = 0x0000ff0000000000
const Rank7 uint64 = 0x00ff000000000000
const Rank8 uint64 = 0xff00000000000000
const NotFileA uint64 = ^FileA
const NotFileH uint64 = ^FileH
const NotRank2 uint64 = ^Rank2
const NotRank4 uint64 = ^Rank4
const NotRank5 uint64 = ^Rank5
const NotRank7 uint64 = ^Rank7

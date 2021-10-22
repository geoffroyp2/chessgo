package bitboard

import "github.com/geoffroyp2/chessgo/engine/constants"

func ShiftN(bb uint64) uint64 {
	return bb << 8
}

func ShiftS(bb uint64) uint64 {
	return bb >> 8
}

func ShiftE(bb uint64) uint64 {
	return bb << 1 & constants.NotFileA
}

func ShiftNE(bb uint64) uint64 {
	return bb << 9 & constants.NotFileA
}

func ShiftSE(bb uint64) uint64 {
	return bb >> 7 & constants.NotFileA
}

func ShiftW(bb uint64) uint64 {
	return bb >> 1 & constants.NotFileH
}

func ShiftNW(bb uint64) uint64 {
	return bb << 7 & constants.NotFileH
}

func ShiftSW(bb uint64) uint64 {
	return bb >> 9 & constants.NotFileH
}
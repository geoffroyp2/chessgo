package bitboard

import "github.com/geoffroyp2/chessgo/engine/constants"

func ShiftN(bb uint64) uint64 {
	return bb << 8
}

func ShiftS(bb uint64) uint64 {
	return bb >> 8
}

func ShiftE(bb uint64) uint64 {
	return bb << 1 & constants.NotAFile
}

func ShiftNE(bb uint64) uint64 {
	return bb << 9 & constants.NotAFile
}

func ShiftSE(bb uint64) uint64 {
	return bb >> 7 & constants.NotAFile
}

func ShiftW(bb uint64) uint64 {
	return bb >> 1 & constants.NotHFile
}

func ShiftNW(bb uint64) uint64 {
	return bb << 7 & constants.NotHFile
}

func ShiftSW(bb uint64) uint64 {
	return bb >> 9 & constants.NotHFile
}
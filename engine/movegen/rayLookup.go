package movegen

var rayLookup = [8][64]uint64{
	{ // N
		0x101010101010100, 0x202020202020200, 0x404040404040400, 0x808080808080800, 0x1010101010101000, 0x2020202020202000, 0x4040404040404000, 0x8080808080808000,
		0x101010101010000, 0x202020202020000, 0x404040404040000, 0x808080808080000, 0x1010101010100000, 0x2020202020200000, 0x4040404040400000, 0x8080808080800000,
		0x101010101000000, 0x202020202000000, 0x404040404000000, 0x808080808000000, 0x1010101010000000, 0x2020202020000000, 0x4040404040000000, 0x8080808080000000,
		0x101010100000000, 0x202020200000000, 0x404040400000000, 0x808080800000000, 0x1010101000000000, 0x2020202000000000, 0x4040404000000000, 0x8080808000000000,
		0x101010000000000, 0x202020000000000, 0x404040000000000, 0x808080000000000, 0x1010100000000000, 0x2020200000000000, 0x4040400000000000, 0x8080800000000000,
		0x101000000000000, 0x202000000000000, 0x404000000000000, 0x808000000000000, 0x1010000000000000, 0x2020000000000000, 0x4040000000000000, 0x8080000000000000,
		0x100000000000000, 0x200000000000000, 0x400000000000000, 0x800000000000000, 0x1000000000000000, 0x2000000000000000, 0x4000000000000000, 0x8000000000000000,
		0, 0, 0, 0, 0, 0, 0, 0,
	}, { // S
		0, 0, 0, 0, 0, 0, 0, 0,
		0x1, 0x2, 0x4, 0x8, 0x10, 0x20, 0x40, 0x80,
		0x101, 0x202, 0x404, 0x808, 0x1010, 0x2020, 0x4040, 0x8080,
		0x10101, 0x20202, 0x40404, 0x80808, 0x101010, 0x202020, 0x404040, 0x808080,
		0x1010101, 0x2020202, 0x4040404, 0x8080808, 0x10101010, 0x20202020, 0x40404040, 0x80808080,
		0x101010101, 0x202020202, 0x404040404, 0x808080808, 0x1010101010, 0x2020202020, 0x4040404040, 0x8080808080,
		0x10101010101, 0x20202020202, 0x40404040404, 0x80808080808, 0x101010101010, 0x202020202020, 0x404040404040, 0x808080808080,
		0x1010101010101, 0x2020202020202, 0x4040404040404, 0x8080808080808, 0x10101010101010, 0x20202020202020, 0x40404040404040, 0x80808080808080,
	}, { // E
		0xFE, 0xFC, 0xF8, 0xF0, 0xE0, 0xC0, 0x80, 0,
		0xFE00, 0xFC00, 0xF800, 0xF000, 0xE000, 0xC000, 0x8000, 0,
		0xFE0000, 0xFC0000, 0xF80000, 0xF00000, 0xE00000, 0xC00000, 0x800000, 0,
		0xFE000000, 0xFC000000, 0xF8000000, 0xF0000000, 0xE0000000, 0xC0000000, 0x80000000, 0,
		0xFE00000000, 0xFC00000000, 0xF800000000, 0xF000000000, 0xE000000000, 0xC000000000, 0x8000000000, 0,
		0xFE0000000000, 0xFC0000000000, 0xF80000000000, 0xF00000000000, 0xE00000000000, 0xC00000000000, 0x800000000000, 0,
		0xFE000000000000, 0xFC000000000000, 0xF8000000000000, 0xF0000000000000, 0xE0000000000000, 0xC0000000000000, 0x80000000000000, 0,
		0xFE00000000000000, 0xFC00000000000000, 0xF800000000000000, 0xF000000000000000, 0xE000000000000000, 0xC000000000000000, 0x8000000000000000, 0,
	}, { // W
		0, 0x1, 0x3, 0x7, 0xF, 0x1F, 0x3F, 0x7F,
		0, 0x100, 0x300, 0x700, 0xF00, 0x1F00, 0x3F00, 0x7F00,
		0, 0x10000, 0x30000, 0x70000, 0xF0000, 0x1F0000, 0x3F0000, 0x7F0000,
		0, 0x1000000, 0x3000000, 0x7000000, 0xF000000, 0x1F000000, 0x3F000000, 0x7F000000,
		0, 0x100000000, 0x300000000, 0x700000000, 0xF00000000, 0x1F00000000, 0x3F00000000, 0x7F00000000,
		0, 0x10000000000, 0x30000000000, 0x70000000000, 0xF0000000000, 0x1F0000000000, 0x3F0000000000, 0x7F0000000000,
		0, 0x1000000000000, 0x3000000000000, 0x7000000000000, 0xF000000000000, 0x1F000000000000, 0x3F000000000000, 0x7F000000000000,
		0, 0x100000000000000, 0x300000000000000, 0x700000000000000, 0xF00000000000000, 0x1F00000000000000, 0x3F00000000000000, 0x7F00000000000000,
	}, { // NE
		0x8040201008040200, 0x80402010080400, 0x804020100800, 0x8040201000, 0x80402000, 0x804000, 0x8000, 0,
		0x4020100804020000, 0x8040201008040000, 0x80402010080000, 0x804020100000, 0x8040200000, 0x80400000, 0x800000, 0,
		0x2010080402000000, 0x4020100804000000, 0x8040201008000000, 0x80402010000000, 0x804020000000, 0x8040000000, 0x80000000, 0,
		0x1008040200000000, 0x2010080400000000, 0x4020100800000000, 0x8040201000000000, 0x80402000000000, 0x804000000000, 0x8000000000, 0,
		0x804020000000000, 0x1008040000000000, 0x2010080000000000, 0x4020100000000000, 0x8040200000000000, 0x80400000000000, 0x800000000000, 0,
		0x402000000000000, 0x804000000000000, 0x1008000000000000, 0x2010000000000000, 0x4020000000000000, 0x8040000000000000, 0x80000000000000, 0,
		0x200000000000000, 0x400000000000000, 0x800000000000000, 0x1000000000000000, 0x2000000000000000, 0x4000000000000000, 0x8000000000000000, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	}, { // NW
		0, 0x100, 0x10200, 0x1020400, 0x102040800, 0x10204081000, 0x1020408102000, 0x102040810204000,
		0, 0x10000, 0x1020000, 0x102040000, 0x10204080000, 0x1020408100000, 0x102040810200000, 0x204081020400000,
		0, 0x1000000, 0x102000000, 0x10204000000, 0x1020408000000, 0x102040810000000, 0x204081020000000, 0x408102040000000,
		0, 0x100000000, 0x10200000000, 0x1020400000000, 0x102040800000000, 0x204081000000000, 0x408102000000000, 0x810204000000000,
		0, 0x10000000000, 0x1020000000000, 0x102040000000000, 0x204080000000000, 0x408100000000000, 0x810200000000000, 0x1020400000000000,
		0, 0x1000000000000, 0x102000000000000, 0x204000000000000, 0x408000000000000, 0x810000000000000, 0x1020000000000000, 0x2040000000000000,
		0, 0x100000000000000, 0x200000000000000, 0x400000000000000, 0x800000000000000, 0x1000000000000000, 0x2000000000000000, 0x4000000000000000,
		0, 0, 0, 0, 0, 0, 0, 0,
	}, { // SE
		0, 0, 0, 0, 0, 0, 0, 0,
		0x2, 0x4, 0x8, 0x10, 0x20, 0x40, 0x80, 0,
		0x204, 0x408, 0x810, 0x1020, 0x2040, 0x4080, 0x8000, 0,
		0x20408, 0x40810, 0x81020, 0x102040, 0x204080, 0x408000, 0x800000, 0,
		0x2040810, 0x4081020, 0x8102040, 0x10204080, 0x20408000, 0x40800000, 0x80000000, 0,
		0x204081020, 0x408102040, 0x810204080, 0x1020408000, 0x2040800000, 0x4080000000, 0x8000000000, 0,
		0x20408102040, 0x40810204080, 0x81020408000, 0x102040800000, 0x204080000000, 0x408000000000, 0x800000000000, 0,
		0x2040810204080, 0x4081020408000, 0x8102040800000, 0x10204080000000, 0x20408000000000, 0x40800000000000, 0x80000000000000, 0,
	}, { // SW
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0x1, 0x2, 0x4, 0x8, 0x10, 0x20, 0x40,
		0, 0x100, 0x201, 0x402, 0x804, 0x1008, 0x2010, 0x4020,
		0, 0x10000, 0x20100, 0x40201, 0x80402, 0x100804, 0x201008, 0x402010,
		0, 0x1000000, 0x2010000, 0x4020100, 0x8040201, 0x10080402, 0x20100804, 0x40201008,
		0, 0x100000000, 0x201000000, 0x402010000, 0x804020100, 0x1008040201, 0x2010080402, 0x4020100804,
		0, 0x10000000000, 0x20100000000, 0x40201000000, 0x80402010000, 0x100804020100, 0x201008040201, 0x402010080402,
		0, 0x1000000000000, 0x2010000000000, 0x4020100000000, 0x8040201000000, 0x10080402010000, 0x20100804020100, 0x40201008040201,
	},
}

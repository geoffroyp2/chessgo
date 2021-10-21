package move

// 24 bits are used to encode: from square, to square, move type, main piece, extra piece (capture)
// 8 bits left for heuristics for later
const (
	fromMask         uint32 = 0x3f
	toOffset         uint32 = 6
	toMask           uint32 = 0x3f << toOffset
	moveTypeOffset   uint32 = 12
	moveTypeMask     uint32 = 0xf << moveTypeOffset
	mainPieceOffset	 uint32 = 16
	mainPieceMask    uint32 = 0xf << mainPieceOffset
	extraPieceOffset uint32 = 20
	extraPieceMask   uint32 = 0xf << extraPieceOffset
)

func CreateMove(from, to, moveType, piece1, piece2 uint32) uint32 {
	return from | (to << toOffset) | (moveType << moveTypeOffset) | (piece1 << mainPieceOffset) | (piece2 << extraPieceOffset)
}

func From(move uint32) uint32 {
	return move & fromMask
}

func To(move uint32) uint32 {
	return (move & toMask) >> toOffset
}

func Type(move uint32) uint32 {
	return (move & moveTypeMask) >> moveTypeOffset
}

func MainPiece(move uint32) uint32 {
	return (move & mainPieceMask) >> mainPieceOffset
}

func ExtraPiece(move uint32) uint32 {
	return (move & extraPieceMask) >> extraPieceOffset
}

func IsCapture(move uint32) bool {
	return Type(move)&CAPTUREMASK != 0
}

func IsPromotion(move uint32) bool {
	return Type(move)&PROMMASK != 0
}
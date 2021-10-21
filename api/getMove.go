package api

import (
	"encoding/json"

	"github.com/geoffroyp2/chessgo/engine/fen"
	"github.com/geoffroyp2/chessgo/engine/move"
	"github.com/geoffroyp2/chessgo/engine/movegen"
)

type MoveResponse struct {
	FEN  string
	Move string
}

func getMove(FENString *string) *[]byte {

	pos := fen.PositionFromFEN(FENString)
	pos.PrintPosition()

	var moveArray [400]uint32
	moveAmount := movegen.GenerateMoves(pos, &moveArray)

	for i := 0; i < moveAmount; i++ {
		move.PrintMove(moveArray[i])
	}

	moveResponse := MoveResponse{FEN: "", Move: ""}
	jsonRsponse, _ := json.Marshal(moveResponse)
	
	return &jsonRsponse
}
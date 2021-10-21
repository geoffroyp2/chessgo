package api

import (
	"encoding/json"

	"github.com/geoffroyp2/chessgo/engine/constants"
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

	var moveArray constants.MoveArray
	moveAmount := movegen.GenerateMoves(pos, &moveArray)

	for i := 0; i < moveAmount; i++ {
		piece := move.MainPiece(moveArray[i]) 
		if (piece == constants.BLACKKNIGHT || piece == constants.WHITEKNIGHT){
			move.PrintMove(moveArray[i])
		}
	}

	moveResponse := MoveResponse{FEN: "", Move: ""}
	jsonRsponse, _ := json.Marshal(moveResponse)
	
	return &jsonRsponse
}
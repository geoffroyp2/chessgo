package api

import (
	"encoding/json"
	"math/rand"

	"github.com/geoffroyp2/chessgo/engine/constants"
	"github.com/geoffroyp2/chessgo/engine/move"
	"github.com/geoffroyp2/chessgo/engine/movegen"
	"github.com/geoffroyp2/chessgo/engine/stringutils/fen"
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

	// for i := 0; i < moveAmount; i++ {
	// 	// piece := move.MainPiece(moveArray[i]) 
	// 	// if (piece == constants.BLACKKNIGHT || piece == constants.WHITEKNIGHT){
	// 	// 	move.PrintMove(moveArray[i])
	// 	// }
	// 	// move.PrintMove(moveArray[i])
	// }
	
	chosenMove := ""

	if moveAmount > 0 {
		randomIdx := rand.Int() % moveAmount
		chosenMove = move.GetMoveString(moveArray[randomIdx])
	}

	moveResponse := MoveResponse{FEN: "", Move: chosenMove}
	jsonRsponse, _ := json.Marshal(moveResponse)
	
	return &jsonRsponse
}
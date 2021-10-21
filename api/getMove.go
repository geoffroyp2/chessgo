package api

import (
	"encoding/json"
)

type MoveResponse struct {
	FEN  string
	Move string
}

func getMove(FENString *string) *[]byte {

	moveResponse := MoveResponse{FEN: "", Move: ""}
	jsonRsponse, _ := json.Marshal(moveResponse)
	
	return &jsonRsponse
}
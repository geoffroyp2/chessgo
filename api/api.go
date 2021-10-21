package api

import (
	"net/http"
)

func EngineHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	FENString := request.URL.Query().Get("FEN")
	if FENString == "" {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Need to provide a FEN string"))
		return
	}

	move := getMove(&FENString)
	writer.Write(*move)
}
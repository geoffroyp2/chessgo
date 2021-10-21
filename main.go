package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/geoffroyp2/chessgo/api"
	"github.com/rs/cors"
)



func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/engine", api.EngineHandler)

	handler := cors.Default().Handler(mux)

	fmt.Println("Listening on port " + port)
	http.ListenAndServe("localhost:" + port, handler)
}
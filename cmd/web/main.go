package main

import (
	"log"
	"net/http"

	"github.com/sunnygolang/chat-websocket/internal/handlers"
)

func main() {
	mux := routes()

	log.Println("Starting channel listenner")
	go handlers.ListenToWsChannel()

	log.Println("Starting web server on port 8080")
	_ = http.ListenAndServe(":8080", mux)
}

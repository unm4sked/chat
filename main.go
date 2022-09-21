package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	wsServer := CrateWebSocketServer()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		FireWsServer(wsServer, w, r)
	})

	if err := http.ListenAndServe(":4000", mux); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

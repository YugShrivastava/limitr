package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Status 	string `json:"status"`
	Body 	string `json:"body"`
}

func endpointHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := Message{
		Status: "Success",
		Body: "You've successfully reached the api",
	}
	
	if err := json.NewEncoder(w).Encode(message); err != nil {
		return
	}
}

func main() {
	http.Handle("/ping", rateLimiter(endpointHandler))
	log.Println("Server listening on port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Println("Error in listening on port 8000:", err)
	}
}
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	Token string `json:"token"`
}

var users = map[string]string{
	"user-identifying-token": "userID1",
	"bonjour":                "aurevoir",
	// Add more tokens and user IDs as needed
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var p Payload
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	userID, ok := users[p.Token]
	if !ok {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if p.Token == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	fmt.Fprint(w, userID)
}

func main() {
	http.HandleFunc("/auth", handler)
	http.ListenAndServe("localhost:8888", nil)
}

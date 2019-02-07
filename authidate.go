package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error")
	} else {
		encoder := json.NewEncoder(w)
		w.Header().Set("Content-Type", "application/json")
		encoder.Encode(user)
	}
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

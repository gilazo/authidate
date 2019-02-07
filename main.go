package main

import (
	"fmt"
	. "gilazo/authidate/auth"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		fmt.Fprintf(w, Password{Value: "123"}.Encrypt())
	default:
		fmt.Fprintf(w, "please post")
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

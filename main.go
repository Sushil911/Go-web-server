package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port localhost:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Error while starting server")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "Request of get")
	case http.MethodPost:
		fmt.Fprintln(w, "Request of post")
	default:
		fmt.Fprintln(w, "Please enter either get or post")
	}
}

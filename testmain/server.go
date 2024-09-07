package server

import (
	"fmt"
	"net/http"
)

func server() {
	// configure path and handler
	http.HandleFunc("/hello", Hello)

	// listen at port 8080 and block main
	fmt.Println("listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Helo world!")
}

package main

import (
	"fmt"
	"net/http"
)

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
}

func main() {
	fmt.Println("Chat App v0.1")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}

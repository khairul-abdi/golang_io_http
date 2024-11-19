package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/example", handler)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated) // membuat status code 201 pada Header
	w.Write([]byte("Status Created"))
}

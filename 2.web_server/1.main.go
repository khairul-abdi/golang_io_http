package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "how are you today!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hi!")
	})

	http.HandleFunc("/index", index)

	http.StatusOK
	fmt.Println("starting web server at localhost:8080")
	http.ListenAndServe(":8080", nil)

}

// localhost:8080 -> 192.168.0.0:8080 -> nginx === https://domain.com

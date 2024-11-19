package main

import (
	"fmt"
	"net/http"
)

func handlerData(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("Menampilkan data"))
	case "POST":
		w.Write([]byte("Menambahkan data"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}

func main() {
	http.HandleFunc("/data", handlerData)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

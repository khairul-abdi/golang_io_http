package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Init Mux dengan menggunakan DefaultServerMux
	mux := http.DefaultServeMux

	// Routing ke actionHandle
	mux.HandleFunc("/action", actionHandle)

	// Membuat handler Middleware yang hanya mengizinkan method GET
	var handler http.Handler = mux
	handler = MiddlewareAllowOnlyGet(handler)

	server := new(http.Server)
	server.Addr = ":8080"
	server.Handler = handler

	fmt.Println("server started at localhost:8080")
	server.ListenAndServe()
}

func MiddlewareAllowOnlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Write([]byte("Only GET is allowed!"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func actionHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Action allowed with GET method"))
}

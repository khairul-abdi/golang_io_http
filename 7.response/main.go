package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/route", handleResponse)
	http.ListenAndServe(":8080", nil)
}

func handleResponse(w http.ResponseWriter, r *http.Request) {
	// Fungsi di bawah digunakan untuk mengatur `HTTP status code` dengan input (statusCode int):
	w.WriteHeader(http.StatusOK)
	// Fungsi di bawah digunakan untuk menentukan `content type` dengan input `Content-Type` dan header value:
	w.Header().Set("Content-Type", "text/plain")
	// Fungsi di bawah digunakan untuk mengatur `response body` dengan input `slice of bytes` yang berisi string:
	w.Write([]byte("Response has been written!"))
}

package main

import (
	"net/http"
)

const USERNAME = "aditira"
const PASSWORD = "1234"

func Login(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte(`something went wrong`))
		return
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		w.Write([]byte(`wrong username/password`))
		return
	}

	w.Write([]byte("Valid authentication"))
}

func main() {
	http.HandleFunc("/login", Login)

	server := new(http.Server)
	server.Addr = ":8080"
	server.ListenAndServe()
}

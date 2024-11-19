package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Data = []student{
	{"Aditira", 22},
	{"Dito", 24},
	{"Ojan", 30},
	{"Tegar", 25},
}

func WithMarshal(w http.ResponseWriter, r *http.Request) {

	jsonInBytes, err := json.Marshal(Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonInBytes)
}

func WithEncode(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func WithDecode(w http.ResponseWriter, r *http.Request) {
	data, err := os.Open("data.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer data.Close()

	var result []student
	err = json.NewDecoder(data).Decode(&result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func WithUnmarshal(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("data.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var result []student
	err = json.Unmarshal(data, &result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func main() {
	//dari json ke struct (ada data yang mau di parse ke struct)
	http.HandleFunc("/marshal", WithMarshal)
	http.HandleFunc("/encode", WithEncode)

	//dari struct ke json (ada data struct yang mau di parse ke json)
	http.HandleFunc("/decode", WithDecode)
	http.HandleFunc("/unmarshal", WithUnmarshal)

	fmt.Println("server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

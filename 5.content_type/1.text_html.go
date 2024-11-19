package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", getTemplate)
	http.ListenAndServe(":8080", nil)
}

func getTemplate(w http.ResponseWriter, r *http.Request) {
	template := `
        <h1>Aditira</h1>
        <p>My Hobby is:</p>
        <ul>
          <li>Programming</li>
          <li>Gaming</li>
        </ul>
    `

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(template))
}

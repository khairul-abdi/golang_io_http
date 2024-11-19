package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type User struct {
	FirstName string
	Age       int
}

const view string = `<html>
    <head>
        <title>Template</title>
    </head>
    <body>
        <h1>Hello</h1>
		{{range $val := .}}
				First Name: {{ .FirstName }}
				Age: {{ .Age }} <br>
            {{else}}
            	Invalid "struct" Users harus berupa array!
        {{end}}
    </body>
</html>`

func main() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {

		var tmpl = template.Must(template.New("main-template").Parse(view))

		if err := tmpl.Execute(w, []User{{"Rogu", 17}, {"Aditira", 21}}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

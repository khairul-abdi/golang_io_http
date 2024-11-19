package main

import (
	"log"
	"os"
	"text/template" //used text template
)

func main() {
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	if err != nil {
		log.Fatal(err)
	}

	err = t.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>")
	if err != nil {
		log.Fatal(err)
	}

}

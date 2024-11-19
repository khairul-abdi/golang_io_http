package main

import (
	"fmt"
	"net/url"
)

var urlString = "https://www.google.com/search?q=golang&hl=id"

func main() {
	var uri, err = url.Parse(urlString)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme=> ", uri.Scheme)     // mendapatkan informasi schema dari URL
	fmt.Println("Host=> ", uri.Host)         // mendapatkan informasi hostname dari URL
	fmt.Println("Path=> ", uri.Path)         // mendapatkan informasi path dari URL
	fmt.Println("RawQuery=> ", uri.RawQuery) // mendapatkan informasi query dari URL

	fmt.Println(uri.Query()["q"][0])  // mendapatkan informasi query 'q'
	fmt.Println(uri.Query()["hl"][0]) // mendapatkan informasi query 'hl'
}

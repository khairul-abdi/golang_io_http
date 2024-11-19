package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/upload", uploadFile)

	// server berjalan di port 8080
	fmt.Println("server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse multipart form, 10 << 20 menentukan maximum upload 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile mengembalikan file pertama untuk key yang diberikan `myFile`
	// ini juga akan mengembalikan FileHeader sehingga kita bisa mendapatkan Filename,
	// Header dan ukuran file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Membuat temporary file di dalam temp-images directory yang mengikuti
	// pola penamaan tertentu
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// baca semua isi file yang kita unggah ke dalam byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// tulis byte array untuk temporary file
	tempFile.Write(fileBytes)
	// mengembalikan bahwa kita telah berhasil mengupload file kita!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

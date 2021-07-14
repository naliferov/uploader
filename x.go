package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {

	// Maximum upload of 2 MB files
	r.ParseMultipartForm(2 << 20)
	fhs := r.MultipartForm.File["uploads"]

	for _, fh := range fhs {
		tempFile, _ := fh.Open()
		newFile, err := os.OpenFile("./"+fh.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}
		io.Copy(newFile, tempFile)
		fmt.Printf("Uploaded File: %+v\n", fh.Filename)
		fmt.Fprintf(w, "Uploaded File: %+v\n", fh.Filename)
	}
}

func display(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./upload.html")
}

func main() {
	http.HandleFunc("/", display)
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}

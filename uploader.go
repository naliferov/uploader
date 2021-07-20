package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {

	//64 mb
	r.ParseMultipartForm(64 << 20)
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

		tempFile.Close()
		newFile.Close()
	}
}

func display(w http.ResponseWriter, r *http.Request) {

	html := `<br>
		<form action="/upload", method="POST" enctype="multipart/form-data">
			<input name="uploads" type="file" multiple="true">
			<input type="submit" value="upload">
		</form>`
	fmt.Fprintf(w, html)
}

func main() {
	http.HandleFunc("/", display)
	http.HandleFunc("/upload", uploadFile)

	fmt.Println("Navigate http://localhost:9999 in your browser")

	http.ListenAndServe(":9999", nil)
}

package main

import (
	"fmt"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)

	fhs := r.MultipartForm.File["myfiles"]
	for _, fh := range fhs {
		f, err := fh.Open()

		defer f.Close()
		//fmt.Printf(f, err)
		// f is one of the files
	}

	// Get handler for filename, size and headers
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
}

func display(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./upload.html")
}

func main() {
	http.HandleFunc("/", display)
	http.HandleFunc("/upload", uploadFile)

	http.ListenAndServe(":8080", nil)
	//cmd.Execute()
}

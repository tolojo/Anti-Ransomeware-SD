package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const uploadPath = "./files"

func main() {
	http.HandleFunc("/upload", uploadFile)
	fs := http.FileServer(http.Dir(uploadPath))
	//se colocado o caminho de apenas /files no browser, vai aparecer os farios ficheiros contidos no servidor
	http.Handle("/files/", http.StripPrefix("/files", fs))

	log.Print("Server started on localhost:8080, use /upload for uploading files and /files/{fileName} for downloading")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading File\n")

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error retrieving file form-data")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("uploading file: %+v\n", handler.Filename)
	fmt.Printf("file size: %+v\n", handler.Size)
	fmt.Printf("MIME header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile("files", "upload-*.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
		return
	}
	tempFile.Write(fileBytes)

	fmt.Fprintf(w, "Sucessfully uploaded file\n")
}

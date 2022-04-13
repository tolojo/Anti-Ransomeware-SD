package main

import (
	"fmt"
	"net/http"
)

type Welcome struct {
	Sale string
	Time string
}

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	file := util.multiPartUpload("ola.txt")
	http.Post("/files", "byte", file)

	fmt.Println(http.ListenAndServe(":80", nil))

}

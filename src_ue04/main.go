package main

import (
	"fmt"
	"net/http"

	"util/packages/util"
)

type Welcome struct {
	Sale string
	Time string
}

func main() {

	http.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
		file := util.MultiPartUpload("ola.txt")
		http.Post("/", "MIME", file)
	})

	fmt.Println(http.ListenAndServe(":80", nil))

}

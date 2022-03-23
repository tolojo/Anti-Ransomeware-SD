package src

import "net/http"

func main() {
	welcome := Welcome{"olá"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
}

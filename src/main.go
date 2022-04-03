package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"example.com/packages/util"
)

type Welcome struct {
	Sale string
	Time string
}

func main() {
	welcome := Welcome{"ola", time.Now().Format(time.Stamp)}
	template := template.Must(template.ParseFiles("template/template.html"))
	comparison := util.Sha256Comparison("ola.txt")
	fmt.Println(comparison)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if sale := r.FormValue("sale"); sale != "" {
			welcome.Sale = sale
		}
		if err := template.ExecuteTemplate(w, "template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
		sha256text := util.Sha256conv("ola.txt")
		fmt.Println(sha256text)
		fmt.Println(w, "You called me!")
	})

	fmt.Println(http.ListenAndServe(":8000", nil))

}

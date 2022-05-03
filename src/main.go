package main

import (
	"encoding/json"
	"errors"
	"example.com/packages/util"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type FileName struct {
	Name string `json:"name"`
}

type Welcome struct {
	Sale string
	Time string
}

type Data struct {
	fileName string
	URL      string
}

func main() {
	welcome := Welcome{"ola", time.Now().Format(time.Stamp)}
	template := template.Must(template.ParseFiles("template/template.html"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if sale := r.FormValue("sale"); sale != "" {
			welcome.Sale = sale
		}
		if err := template.ExecuteTemplate(w, "template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/save", func(w http.ResponseWriter, response *http.Request) {

		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}

		response.Body.Close()
		fmt.Println(string(bytes))
		var fileResponse FileName
		errUnmarshal := json.Unmarshal(bytes, &fileResponse)

		if errUnmarshal != nil {
			log.Fatal(errUnmarshal)
		}
		log.Printf("%+v", fileResponse)

		data := &Data{
			fileName: fileResponse.Name,
			URL:      "http://10.72.182.207/files/" + fileResponse.Name,
		}

		data.download("securityCopy/")

		sha256text := util.Sha256conv(fileResponse.Name)
		fmt.Println(sha256text)
	})

	http.HandleFunc("/validate", func(w http.ResponseWriter, response *http.Request) {
		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}

		response.Body.Close()
		fmt.Println(string(bytes))
		var fileResponse FileName
		errUnmarshal := json.Unmarshal(bytes, &fileResponse)

		if errUnmarshal != nil {
			log.Fatal(errUnmarshal)
		}
		log.Printf("%+v", fileResponse)

		data := &Data{
			fileName: fileResponse.Name,
			URL:      "http://10.72.182.207/files/" + fileResponse.Name,
		}

		data.download("temp/")

		comparison := util.Sha256Comparison(fileResponse.Name)
		fmt.Println(comparison)
		e := os.Remove("temp/" + fileResponse.Name)
		if e != nil {

		}
	})

	fmt.Println(http.ListenAndServe(":8000", nil))

}

func (data *Data) download(Dir string) error {
	response, err := http.Get(data.URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response status")
	}

	file, err := os.Create(Dir + data.fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)

	if err != nil {
		return err
	}

	return nil
}

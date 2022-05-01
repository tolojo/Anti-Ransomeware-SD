package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type FileName struct {
	Name string `json:"name"`
}

type Data struct {
	fileName string
	URL      string
}

func main() {
	//comparison := util.Sha256Comparison("ola.txt")
	//fmt.Println(comparison)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

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
			URL:      "http://192.168.1.116/files/" + fileResponse.Name,
		}

		data.download()

		//sha256text := util.Sha256conv("abc.txt")
		//fmt.Println(sha256text)
	})

	http.HandleFunc("/receive", func(w http.ResponseWriter, r *http.Request) {

	})

	fmt.Println(http.ListenAndServe(":8000", nil))

}

func (data *Data) download() error {
	response, err := http.Get(data.URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response status")
	}

	file, err := os.Create("securityCopy/" + data.fileName)

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

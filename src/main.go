package main

import (
	bytes2 "bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"example.com/packages/util"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
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
	ipServerPub := "https://10.72.231.72:8443"
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
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
			URL:      ipServerPub + "/files/" + fileResponse.Name,
		}

		log.Printf("%+v", data.download("securityCopy/"))

		sha256text := util.Sha256conv(fileResponse.Name)
		fmt.Println(sha256text)
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/validate", func(w http.ResponseWriter, response *http.Request) {
		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}

		response.Body.Close()
		log.Println(string(bytes))
		var fileResponse FileName
		errUnmarshal := json.Unmarshal(bytes, &fileResponse)

		if errUnmarshal != nil {
			log.Fatal(errUnmarshal)
		}
		log.Printf("%+v", fileResponse)

		data := &Data{
			fileName: fileResponse.Name,
			URL:      ipServerPub + "/files/" + fileResponse.Name,
		}

		data.download("temp/")

		comparison := util.Sha256Comparison(fileResponse.Name)
		log.Printf("%+v", comparison)
		e := os.Remove("temp/" + fileResponse.Name)
		if e != nil {
		}
		if comparison == false {
			body := &bytes2.Buffer{}
			writer := multipart.NewWriter(body)
			fw, err := writer.CreateFormFile("myFile", fileResponse.Name)
			if err != nil {
				log.Fatal(err)
			}

			myFile, err := os.Open("securityCopy/" + fileResponse.Name)
			if err != nil {
				log.Fatal(err)
			}
			_, err = io.Copy(fw, myFile)
			writer.Close()

			req, err := http.NewRequest("POST", ipServerPub+"/upload", bytes2.NewReader(body.Bytes()))
			req.Header.Set("Content-Type", writer.FormDataContentType())
			client := &http.Client{}
			rsp, _ := client.Do(req)
			if rsp.StatusCode != http.StatusOK {
				log.Printf("Request Failed with response: %d", rsp.StatusCode)
			}
		}

	})

	fmt.Println(http.ListenAndServe(":8000", nil))

}

func (data *Data) download(Dir string) error {
	fmt.Println("abc")
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

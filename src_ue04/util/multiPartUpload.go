package util

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
)

func MultiPartUpload(fname string) *bytes.Buffer {
	fileDir, _ := os.Getwd()
	fileName := fname
	filePath := path.Join(fileDir, fileName)

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	return body

}

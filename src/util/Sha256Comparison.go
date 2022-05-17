package util

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"
)

func Sha256Comparison(s string) bool {
	var securityDirectory = "temp/"
	//abrir o ficheiro s e passa-lo para a variavel f
	f, err := os.Open(securityDirectory + s)
	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	//criar uma nova hash e atribui-la ao ficheiro f
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	//Passar a hash para string e guardar na variavel xa
	xa := hex.EncodeToString(h.Sum(nil))

	file, err := os.Open("hash/hash_" + s)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	newXa := ""

	for _, eachLn := range text {
		if eachLn != "" {
			newXa = eachLn
		}
	}
	log.Print(xa)
	log.Print(newXa)

	return xa == newXa

}

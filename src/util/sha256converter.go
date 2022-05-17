package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func Sha256conv(s string) int {
	//abrir o ficheiro s e passa-lo para a variavel f
	f, err := os.Open("securityCopy/" + s)
	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	//criar uma nova hash e atribui-la ao ficheiro f
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	//Passar a hash para string e guardar na variavel xa
	xa := hex.EncodeToString(h.Sum(nil))
	log.Print(xa)
	//criar ficheiro com o nome "hash_" + nome do parametro s
	a := os.WriteFile("hash/hash_"+s, []byte(xa), 0755)

	fmt.Printf("Ficheiro escrito com %+v bytes", a)
	return 0
}

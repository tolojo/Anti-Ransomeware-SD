package sha256conv

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func sha256conv(caminho string) int {
	f, err := os.Open(caminho)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x", h.Sum(nil))
	return 0
}

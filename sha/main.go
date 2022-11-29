package main

import (
	"crypto/sha1"
	"encoding/hex"
	"log"
)

func main() {
	bytes := []byte("nick♥")
	hasher := sha1.New()
	n, err := hasher.Write(bytes)
	if err != nil {
		log.Fatal("Chould not generate hash", err)
	}

	log.Println("number of bytes written", n)

	sum := hasher.Sum(nil)
	result := hex.EncodeToString(sum)
	log.Println(result)
}

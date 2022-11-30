package main

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"log"
	"os"
)

func main() {
	// Opening the file
	file, err := os.Open("testFile.txt")
	if err != nil {
		log.Fatal("Could not open file", err)
	}

	// Defer closing the file, but also log any error in doing so.
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal("Problem closing the file", err)
		}
	}()

	hasher := sha1.New()

	// The same buffer is re-used to avoid memory allocations.
	buffer := make([]byte, 4*1024)
	for {
		numBytes, err := io.CopyBuffer(hasher, file, buffer)
		if err == io.EOF {
			break
		}

		if numBytes == 0 {
			break
		}

		if err != nil {
			log.Fatal("Problem copying buffer", err)
		}

	}

	sum := hasher.Sum(nil)
	result := hex.EncodeToString(sum)
	log.Println(result)
}

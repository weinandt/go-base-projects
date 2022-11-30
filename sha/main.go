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
		// This could probably be done more efficiently.
		// While writing the bytes and computing the hash, we could be going back to the disk to read more information.
		numBytes, err := io.CopyBuffer(hasher, file, buffer)
		if err != nil {
			// Don't need to check for EOF.
			// From golang docs: "A successful Copy returns err == nil, not err == EOF. Because Copy is defined to read from src until EOF, it does not treat an EOF from Read as an error to be reported."
			log.Fatal("Problem copying buffer", err)
		}

		if numBytes == 0 {
			break
		}
	}

	sum := hasher.Sum(nil)
	result := hex.EncodeToString(sum)
	log.Println(result)
}

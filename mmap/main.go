package main

import (
	"log"
	"os"
	"syscall"
)

func main() {
	// AWS supports files at tmp directory: https://aws.amazon.com/blogs/aws/aws-lambda-now-supports-up-to-10-gb-ephemeral-storage/
	filePath := "/tmp/myFile"
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal("Problem opening file.", err)
	}

	// Todo: figure out this call.
	bytes, err := syscall.Mmap(int(f.Fd()), 0, 10000, syscall.PROT_WRITE, syscall.MAP_PRIVATE)
	if err != nil {
		log.Fatal("Problem with mmap", err)
	}

	log.Println("length: ", len(bytes))

	// Not using defer so we can log if there is an error closing the file.
	if err = f.Close(); err != nil {
		log.Fatal("Problem closing file", err)
	}
}

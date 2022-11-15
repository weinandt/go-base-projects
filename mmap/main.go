package main

import (
	"log"
	"syscall"
	"time"
)

func main() {
	// TODO: Why is map anon required?
	// TODO: why is a file unecessary?
	bytes, err := syscall.Mmap(-1, 0, 1000000, syscall.PROT_WRITE|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_SHARED)
	if err != nil {
		log.Fatal("Problem with mmap", err)
	}

	log.Println("length: ", len(bytes))
	log.Println(bytes[0])

	time.Sleep(time.Second * 10)
}

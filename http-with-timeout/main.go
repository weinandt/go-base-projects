package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

func makeRequest() (string, error) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		return "", err
	}

	if res.StatusCode != 200 {
		return "", errors.New("Response was not 200")
	}

	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	bodyString := string(bodyBytes)

	return bodyString, nil
}

func makeCallOrTimeOutUsingTimePackage(timeToWaitInMs uint) (string, error) {
	var res string
	var err error
	done := make(chan struct{})
	go func() {
		res, err = makeRequest()
		close(done)
	}()

	select {
	case <-done:
		return res, err
	case <-time.After(time.Millisecond * time.Duration(timeToWaitInMs)):
		return "", errors.New("Timed out")
	}
}

func makeCallOrTimeoutManualContextCancellation(timeToWaitInMs uint) (string, error) {
	cxtTimeout, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(timeToWaitInMs))

	// Cancel must be called to avoid leaking reasources.
	defer cancel()

	// Becuase it is possible the request times out before the channel is written to,
	// we need to make sure writing to the channel after the context is timedout does not block the go routine.
	// Could also use a buffered channel of size 1 here.
	// as writing to a buffered channel will not block if the buffer is not full.
	var res string
	var err error
	done := make(chan struct{})
	go func() {
		res, err = makeRequest()
		close(done)
	}()

	select {
	case <-done:
		return res, err
	case <-cxtTimeout.Done():
		return "", errors.New("Timed out")
	}
}

func main() {
	res, err := makeCallOrTimeoutManualContextCancellation(1000)
	if err != nil {
		log.Fatal("Problem making request", err)
	}

	log.Println(res)
}

package main

import (
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

func main() {
	res, err := makeCallOrTimeOutUsingTimePackage(1000)
	if err != nil {
		log.Fatal("Problem making request", err)
	}

	log.Println(res)
}

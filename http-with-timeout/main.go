package main

import (
	"errors"
	"io"
	"log"
	"net/http"
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

func main() {
	res, err := makeRequest()
	if err != nil {
		log.Fatal("Problem making request", err)
	}

	log.Println(res)
}

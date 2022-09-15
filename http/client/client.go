package main

import (
	"fmt"
	"net/http"
)

func makeRequest() {
	res, err := http.Get("http://localhost:50000/user")
	if err != nil {
		fmt.Println("Problem making request.")
	}

	if res.StatusCode != 200 {
		fmt.Println("Did not return 200.")
	}
}

func main() {
	makeRequest()
}

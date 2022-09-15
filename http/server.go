package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	myUser := struct {
		Name string `json:"name"`
	}{
		Name: "Nick",
	}

	response, _ := json.Marshal(myUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", getUser).Methods("GET")
	http.ListenAndServe(":50000", router)
}

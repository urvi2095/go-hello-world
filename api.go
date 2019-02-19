package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	time2 "time"
)

type Time struct {
	Now string `json:"now,omitempty"`
	Message string `json:"m"`
}

func GetTime(w http.ResponseWriter, r *http.Request) {

	time := Time{
		Now: time2.Now().String(),
		Message: "Hello" ,
	}

	log.Printf("Get request from %v", r.RemoteAddr)

	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(time)
}

func CreateTime(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	fmt.Printf("%v", params)

	var time Time
	json.NewDecoder(r.Body).Decode(&time)
	fmt.Printf("%v\n", time)

	//for index, value := range 2 {
	//	fmt.Printf("%v %v\n", index, value)
	//}


}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetTime).Methods("GET")
	router.HandleFunc("/", CreateTime).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//ToDo struct
type ToDo struct {
	Item        string    `json:"Item"`
	Description string    `json:"Description"`
	Complete    bool      `json:"Complete"`
	DueDate     time.Time `json:"DueDate"`
}

//TodoItems array of things to do
type TodoItems []ToDo

func getAllItems(w http.ResponseWriter, r *http.Request) {
	todoitems := TodoItems{
		ToDo{Item: "GO Rest API", Description: "Make a Go Rest API you lazy manchild", Complete: false, DueDate: time.Now()},
	}
	fmt.Println("Todo items")
	json.NewEncoder(w).Encode(todoitems)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/todoitems", getAllItems)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}

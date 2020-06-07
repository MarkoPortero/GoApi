package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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

func postTodoItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post worked bro")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	//mux allows much easier verb usage i.e. Methods("GET")
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/home", homePage)
	myRouter.HandleFunc("/todoitems", getAllItems).Methods("GET")
	myRouter.HandleFunc("/todoitems", postTodoItems).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}

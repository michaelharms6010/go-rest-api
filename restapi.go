package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
)

type Quote struct {
	Id string `json:"Id"`
	Content string `json:"quote"`
	Speaker string `json:"speaker"`
}

var Quotes []Quote

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
	fmt.Println("Homepage Endpoint hit")
}



func returnAllQuotes(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Quotes)
}

func createNewQuote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: Create Quote")
    reqBody, _ := ioutil.ReadAll(r.Body)
    var quote Quote 
    json.Unmarshal(reqBody, &quote)
    Quotes = append(Quotes, quote)

    json.NewEncoder(w).Encode(quote)
}

func returnSingleQuote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleQuote")
    vars := mux.Vars(r)
    id := vars["id"]

    for _, quote := range Quotes {
        if quote.Id == id {
            json.NewEncoder(w).Encode(quote)
        }
    }
}
func updateQuote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateQuote")
	reqBody, _ := ioutil.ReadAll(r.Body)
    var quote Quote 
    json.Unmarshal(reqBody, &quote)
}


func deleteQuote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteQuote")
    vars := mux.Vars(r)
    id := vars["id"]

    for index, quote := range Quotes {
        if quote.Id == id {
            Quotes = append(Quotes[:index], Quotes[index+1:]...)
        }
    }

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/quotes", returnAllQuotes)
	myRouter.HandleFunc("/quote", createNewQuote).Methods("POST")
	myRouter.HandleFunc("/quote/{id}", deleteQuote).Methods("DELETE")
	myRouter.HandleFunc("/quote/{id}", updateQuote).Methods("PUT")
	myRouter.HandleFunc("/quote/{id}", returnSingleQuote)
	
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {
	Quotes = []Quote{
		Quote{Id: "1", Content: "You would make a ship sail against the winds and currents by lighting a bon-fire under her deck? I have no time for such nonsense.", Speaker: "Napoleon, on Robert Fulton's Steamship"},
		Quote{Id: "2", Content: "Never trust a computer you can't throw out a window.", Speaker: "Steve Wozniak"},
	}
	
	handleRequests()
}
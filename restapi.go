package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Quote struct {
	Content string `json:"quote"`
	Speaker string `json:"speaker"`
}

var Quotes []Quote

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
	fmt.Println("Homepage Endpoint hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/quotes", returnAllQuotes)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func returnAllQuotes(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Quotes)
}

func main() {
	Quotes = []Quote{
		Quote{Content: "You would make a ship sail against the winds and currents by lighting a bon-fire under her deck? I have no time for such nonsense.", Speaker: "Napoleon, on Robert Fulton's Steamship"},
		Quote{Content: "Never trust a computer you can't throw out a window.", Speaker: "Steve Wozniak"},
	}
	
	handleRequests()
}
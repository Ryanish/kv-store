// Create Kv Store - start creating a restful API using HTTP than can obtain data from the kv-store which is basically a map

package main

import (
	//"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World")
	startDatabase()
	webhttp()
	// Define structure
	// var CRUDresponse = []APIresponse{
	// 	APIresponse{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	// 	APIresponse{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	// }
	handleRequests()
}

// Create KV Store - essentially a "map" in Go. TODO. Needs to be set to CRUD in memory, rather than be hardcoded
func startDatabase() {
	kvstore := make(map[string]int)

	kvstore["entry one"] = 1
	kvstore["entry two"] = 2
	kvstore["entry three"] = 3

	fmt.Println("Database Ready")
}

////////////////////////////////////////

// main list of endpoints
func webhttp() {
	http.HandleFunc("/kvstore", primaryPage)
	http.HandleFunc("/returnalltitles", returnAllTitles)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

// HTTP web services - "viewhandler is too vague"
func primaryPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint Hit: kvstore")
	message := []byte("KV Store Service is running - HOORAY FOR GOPHERS.")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func returnAllTitles(writer http.ResponseWriter, r *http.Request) {
	//loads in command once hit - TODO - could log this out?
	fmt.Println("Endpoint Hit: returnAllArticles")
	message := []byte("this page loads all the articles titles.")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
// 	json.NewEncoder(w).Encode(CRUDresponse)
}

/////////////////////////

type CRUDresponse struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func handleRequests() {
//http.HandleFunc("/", homePage)
// add our articles route and map it to our 
// returnAllArticles function like so
//http.HandleFunc("/articles", returnAllArticles)
log.Fatal(http.ListenAndServe(":10000", nil))
}
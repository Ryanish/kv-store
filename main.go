// Create Kv Store - start creating a restful API using HTTP than can obtain data from the kv-store which is basically a map

//Develop an in-memory key value store. The store must support data storage and retrieval through TCP, UDP, HTTP and HTTPS REST interfaces. The store must cater for arbitrary data types.

// good source - https://earthly.dev/blog/golang-http/

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"encoding/json"
)

func main() {
	//Endpoints
	http.HandleFunc("/kvstore_get", handlerGet)

	//Basic Web Server
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handlerGet(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/kvstore_get" {
		http.Error(writer, "404 not found", http.StatusNotFound)
		return
	}

	 if request.Method != "GET" {
		http.Error(writer, "Method is not supported.", http.StatusNotFound)
	 }

	// KV Store
	kvstore := make(map[string]int)
	//adding in 2 kv pairs for example / hard coded purposes
	kvstore["ryan"] = 36
	kvstore["mikey"] =  29

	// adds a kv pair - I don't care for handling conflicts yet


	// deletes a kv pair 
	//delete(kvstore["mikey"])

	fmt.Println(kvstore["ryan"])
}


// Post - need to confirm GET first, get the basics nailed and documented.

// func handlerEncode(writer http.ResponseWriter, request *http.Request) {
// 	if request.URL.Path != "/kvstoreencode" {
// 		http.Error(writer, "404 not found", http.StatusNotFound)
// 		return
// 	}

// 	// if request.Method != "GET" {
// 	// 	http.Error(writer, "Method is not supported.", http.StatusNotFound)
// 	// }

// 	fmt.Println("Endpoint Hit: kvstoreEncode")
// 	JohnSmith := Kvstore{
// 		ID: 3,
// 		DataEntry: "John",
// 	}

// 	json.NewEncoder(writer).Encode(JohnSmith)

// }

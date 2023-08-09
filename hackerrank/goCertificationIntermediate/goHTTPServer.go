// Implement an HTTP server that has three routes and maintains a database of
// the worlds largests lakes.
// the first route runs a handler postHandler which accepts a POST request with JSON-encoded lake
// information and posts it to the database.

// {
// 	"type": "post",
//  "payload": {
// "id": "id00",
// "name": "Caspian Sea",
// "area": 371000
// 	}
// }

// the second route runs a handler deleteHandler which deletes the lake form
// the database by id. If the payload is not present in the dataset, the server
// returns a 404 response

// {
// 	"type": "delete",
// 	"payload": {
// 		"id": "id00"
// 	}
// }

// the third route runs a handler getHandler whiuch takes a lake from the database by id and returns it
// to the caller which prints the name and the area of the lake.
// If the id is not found in the database, the server returns a string message
// "404 not found".
// Otherwise, it returns the payload object corresponding to the id.
// {
// 	"type": "get",
// 	"payload": {
// 		"id": "id00"
// 	}
// }

// implement the server which will be queried by the program and print the output to STDOUT
// Note: the program uses these structs:
// type Lake struct {

// 	Name string
// 	Area int32
// }

package main

import "net/http"

func postHandler(w http.ResponseWriter, r *http.Request) {
	// parse the request body
	// create a new lake object
	// add the lake to the database
	// return a 200 response
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	// parse the request body
	// delete the lake from the database
	// return a 200 response
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	// parse the request body
	// get the lake from the database
	// return a 200 response
}

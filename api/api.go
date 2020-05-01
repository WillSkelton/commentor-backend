package api

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port         = "42201"
	tempHomePage = `
	<h1>Welcome to the Commentor Homepage</h1>
	<p>There isn't much here now but there will be soon.</p>
	`
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: /homepage")
	fmt.Fprintf(w, tempHomePage)
}

func handleRequests() {
	// This tells the server to run the function homePage when a request is sent to
	// "/" which is the home page
	http.HandleFunc("/", homePage)

	fmt.Printf("Listening on port: %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// Start will setup the routes and their respective functions as well as telling the
// server which port to listen on
func Start() {
	handleRequests()
}

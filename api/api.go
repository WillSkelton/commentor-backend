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
	http.HandleFunc("/", homePage)

	fmt.Println("Listening on port: 42201")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func Start() {
	handleRequests()
}

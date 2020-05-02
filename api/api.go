package api

import (
	"commentor-backend/lib/driver"
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

var (
	singleton *driver.Driver
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: /homepage")
	enableCors(&w)
	fmt.Fprintf(w, tempHomePage)
}

func openDirectory(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	keys, ok := r.URL.Query()["wd"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'wd' is missing")
		return
	}

	wd := keys[0]

	fmt.Println("Url Param 'wd' is: " + string(wd))

	var err error
	if singleton, err = driver.NewDriver(wd); err != nil {
		log.Fatal(err)
		return
	}

	// for _, value := range singleton.FileManager {
	// 	fmt.Println(value)
	// }

}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/opendirectory", openDirectory)

	fmt.Printf("Listening on port: %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

// Start will setup the routes and their respective functions as well as telling the
// server which port to listen on
func Start() {
	handleRequests()
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

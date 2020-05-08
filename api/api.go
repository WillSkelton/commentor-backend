package api

import (
	"commentor-backend/lib/driver"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	port         = "42201"
	tempHomePage = `
	<h1>Welcome to the Commentor Homepage</h1>
	<p>There isn't much here now but there will be soon.</p>
	`
)

var (
	// singleton *driver.Driver
	singleton = driver.NewDriver("")
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
		http.Error(w, "Url Param 'wd' is missing", 400)
		return
	}

	wd := keys[0]

	singleton = driver.NewDriver(wd)

	var err error
	if _, err = os.Stat(wd); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	// fmt.Println("Working Directory is: " + string(wd))

	// if singleton = driver.NewDriver(wd); err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }

	if err = singleton.GatherFiles(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var res []byte
	if res, err = json.Marshal(singleton.FileManager); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// fmt.Println(string(res))
	// fmt.Println(singleton.FileManager[0].Functions[63])
	fmt.Fprintf(w, "%v", (string(res)))

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

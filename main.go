package main

import (
	"fmt"
	"log"
	"net/http"
	"nextnotes/notes"
	"os"
)

func main() {
	http.HandleFunc("/notes", notes.Notes)

	var port string
	var isset bool
	port, isset = os.LookupEnv("PORT")

	if !isset {
		port = "8001"
	}

	fmt.Printf("Backend init at localhost: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

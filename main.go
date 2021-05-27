package main

import (
	"fmt"
	"log"
	"net/http"
	"nextnotes/notes"
	"nextnotes/notifications"
	"os"
)

func main() {
	http.HandleFunc("/notes", notes.Notes)
	http.HandleFunc("/notifications", notifications.Notifications)

	var port string
	var isset bool
	port, isset = os.LookupEnv("PORT")

	if !isset {
		port = "8001"
	}

	fmt.Printf("Backend init at localhost: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"nextnotes/notes"
)

func main() {
	http.HandleFunc("/notes", notes.Notes)
	fmt.Println("Backend init at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./client")))

	http.HandleFunc("/api/get-preview", PreviewHandler)
	
	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

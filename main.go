package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePage)

	fmt.Printf("Server listening on port 8080")
	log.Panic(http.ListenAndServe(":8080", nil))
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "Welcome to the inspirational quote API homepage")
	if err != nil {
		log.Panic(err)
	}
}
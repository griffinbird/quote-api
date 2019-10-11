package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type quoteStruct struct {
	Quote string `json:"quote"`
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/quotes", quotes)

	fmt.Printf("Server listening on port 8080")
	log.Panic(http.ListenAndServe(":8080", nil))
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	writeResponseOrPanic(writer, "Welcome to the inspirational quote API homepage")
}

func quotes(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		newQuote(writer, request)
	} else {
		writeResponseOrPanic(writer, "Invalid request method")
	}
}

func newQuote(writer http.ResponseWriter, request *http.Request) {
	quote, err := NewQuoteFromRequest(request)
	if err != nil {
		writeResponseOrPanic(writer, fmt.Sprintf("error unable to create a quote from request data.\nmessage: %s\n", err.Error()))
		return
	}

	writeResponseOrPanic(writer, fmt.Sprintf("Quote added: \"%s\"\n", quote.Quote))
}

// Will write a response using the http.ResponseWriter. If it fails it will panic.
func writeResponseOrPanic(writer http.ResponseWriter, message string) {
	_, err := fmt.Fprint(writer, message)
	if err != nil {
		log.Panic(err)
	}
}
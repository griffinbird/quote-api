package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type JsonMessage struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/quotes", quotes)

	fmt.Printf("Server listening on port 8080")
	log.Panic(http.ListenAndServe(":8080", nil))
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	writeJson(writer, &JsonMessage{"welcome to the inspirational quote homepage"}, 200)
}

func quotes(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		newQuote(writer, request)
		return
	}

	if request.Method == http.MethodGet {
		getRandomQuote(writer)
		return
	}

	writeJson(writer, &JsonMessage{"invalid request method"}, 422)
}

func newQuote(writer http.ResponseWriter, request *http.Request) {
	quote, err := NewQuoteFromRequest(request)
	fmt.Printf("%+v\n", quote)
	if err != nil {
		writeJson(writer, &JsonMessage{err.Error()}, 422)
		return
	}

	err = quote.storeInDatabase()
	if err != nil {
		writeJson(writer, &JsonMessage{err.Error()}, 422)
		return
	}

	writeJson(writer, &JsonMessage{"quote added"}, 200)
}

func getRandomQuote(writer http.ResponseWriter) {
	quoteStruct, err := RandomQuoteFromDatabase()
	if err != nil {
		writeJson(writer, JsonMessage{err.Error()}, 422)
		return
	}
	writeJson(writer, quoteStruct, 200)
}

func writeJson(writer http.ResponseWriter, data interface{}, status int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(jsonBytes)) //logs quote displayed from database
	writeResponseOrPanic(writer, string(jsonBytes))
}

// Will write a response using the http.ResponseWriter. If it fails it will panic.
func writeResponseOrPanic(writer http.ResponseWriter, message string) {
	_, err := fmt.Fprint(writer, message)
	if err != nil {
		log.Panic(err)
	}
}
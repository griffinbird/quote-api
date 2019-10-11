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
	_, err := fmt.Fprintf(writer, "Welcome to the inspirational quote API homepage")
	if err != nil {
		log.Panic(err)
	}
}

func quotes(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		newQuote(writer, request)
	} else {
		_, err := fmt.Fprintf(writer, "Invalid request method")
		if err != nil {
			log.Panic(err)
		}
	}
}

func newQuote(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		_, err = fmt.Fprintf(writer, "Error: cannot read request body.")
		if err != nil {
			log.Panic(err)
		}
		return
	}

	var quote quoteStruct
	err = json.Unmarshal(body, &quote)
	if  err != nil {
		_, err = fmt.Fprintf(writer, "Error: cannot unmarshal request body.")
		if err != nil {
			log.Panic(err)
		}
		return
	}

	if quote.Quote == "" {
		_, err = fmt.Fprintf(writer, "Error: please enter a quote.")
		if err != nil {
			log.Panic(err)
		}
		return
	}

	_, err = fmt.Fprintf(writer, "Quote added: \"%s\"\n", quote.Quote)
	if err != nil {
		log.Panic(err)
	}
}
package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type QuoteStruct struct {
	Quote string `json:"quote"`
}

// Create a new quote from an HTTP request.
func NewQuoteFromRequest(request *http.Request) (*QuoteStruct, error) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	var quote QuoteStruct
	err = json.Unmarshal(body, &quote)
	if err != nil {
		return nil, err
	}

	if quote.Quote == "" {
		return nil, errors.New("the quote cannot be empty")
	}

	return &quote, nil
}

func (quote *QuoteStruct) storeInDatabase() error {
	query := "INSERT INTO quotes (id, quote) VALUES (?, ?)"
	_, err := ExecDB(query, nil, quote.Quote)

	return err
}
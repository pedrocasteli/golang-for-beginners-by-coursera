package main

import (
	"errors"
	"fmt"
	"strings"
)

func getBookFromString(stringToLookUp string) (string, error) {
	for _, book := range books {
		if strings.Contains(strings.ToUpper(book), strings.ToUpper(stringToLookUp)) {
			return book, nil
		}
	}

	message := fmt.Sprintf("Couldn't find a book title with '%s'", stringToLookUp)

	return "", errors.New(message)
}

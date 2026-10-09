package main

import "fmt"

func main() {
	var userInput, book string
	var bookError error

	validBookEntered := false

	for !validBookEntered {
		fmt.Print("Search a book title: ")
		fmt.Scan(&userInput)

		book, bookError = getBookFromString(userInput)

		if bookError == nil {
			fmt.Printf("Book: %s\n", book)
			validBookEntered = true
		} else {
			fmt.Println(bookError)
		}
	}

}

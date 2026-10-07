package main

import "fmt"

func main() {

	var userName string

	fmt.Println("What is your name?")
	fmt.Scanln(&userName)
	fmt.Println("Welcome, " + userName + "!")

}

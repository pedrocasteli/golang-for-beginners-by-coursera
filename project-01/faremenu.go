package main

import "fmt"

func main() {
	var origin, destination, cabinClass string
	var originCity City
	var originError error
	var validOriginEnterned bool = false

	for !validOriginEnterned {
		fmt.Println("*** Silverkarma Airfare Calculator ***")
		fmt.Print("Enter origin code: ")
		fmt.Scanln(&origin)

		originCity, originError = getCityFromCode(origin)

		if originError == nil {
			fmt.Println("You've entered " + originCity.cityName)
			validOriginEnterned = true
		} else {
			fmt.Println(originError)
		}
	}

	fmt.Print("Enter destination code: ")
	fmt.Scanln(&destination)

	fmt.Print("Enter cabin class code: ")
	fmt.Scanln(&cabinClass)

	fmt.Println(cities[0].cityName)
	fmt.Println(cities[1].latitude)
	fmt.Println()
}

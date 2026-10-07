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

	var destinationCity City
	var destinationError error
	var validDestinationEnterned bool = false

	for !validDestinationEnterned {
		fmt.Print("Enter destination code: ")
		fmt.Scanln(&destination)

		destinationCity, destinationError = getCityFromCode(destination)

		if destinationError == nil {
			fmt.Println("You've entered " + destinationCity.cityName)
			validDestinationEnterned = true
		} else {
			fmt.Println(destinationError)
		}
	}

	validCabinClassEntered := false
	var enteredCabinClass CabinClass
	var enteredCabinClassErr error

	for !validCabinClassEntered {

		fmt.Print("Enter cabin class code: ")
		fmt.Scanln(&cabinClass)

		enteredCabinClass, enteredCabinClassErr = getCabinClassFromCode(cabinClass)

		if enteredCabinClassErr == nil {
			fmt.Println("You've entered " + enteredCabinClass.className + " class")
			validCabinClassEntered = true
		} else {
			fmt.Println(enteredCabinClassErr)
		}
	}
}

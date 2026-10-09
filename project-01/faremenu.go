package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

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
			fmt.Println("You've entered: " + originCity.cityName)
			validOriginEnterned = true
		} else {
			fmt.Println(originError)
		}
	}

	var destinationCity City
	var destinationError error
	validDestinationEnterned := false

	for !validDestinationEnterned {
		fmt.Print("Enter destination code: ")
		fmt.Scanln(&destination)

		destinationCity, destinationError = getCityFromCode(destination)

		if destinationError == nil {
			fmt.Println("You've entered: " + destinationCity.cityName)
			validDestinationEnterned = true
		} else {
			fmt.Println(destinationError)
		}
	}

	validCabinClassEntered := false
	var enteredCabinClass CabinClass
	var enteredCabinClassErr error

	for !validCabinClassEntered {

		fmt.Println("Available classes: ")

		for _, cabinInfo := range cabinClasses {
			fmt.Printf("%s: %s\n", cabinInfo.className, cabinInfo.code)
		}

		fmt.Print("Enter a class code: ")
		fmt.Scanln(&cabinClass)

		enteredCabinClass, enteredCabinClassErr = getCabinClassFromCode(cabinClass)

		if enteredCabinClassErr == nil {
			fmt.Println("You've entered: " + enteredCabinClass.className + " class")
			validCabinClassEntered = true
		} else {
			fmt.Println(enteredCabinClassErr)
		}
	}

	distance := CalculateDistance(
		float64(destinationCity.longitude)/10000,
		float64(destinationCity.latitude)/10000,
		float64(originCity.longitude)/10000,
		float64(originCity.latitude)/10000)

	fmt.Printf("\nDistance: %.1f km\n", distance)
	rate := decimal.New(int64(enteredCabinClass.rate), -2)

	fmt.Printf("$ per km = %s\n", rate.StringFixed(2))

	fare := decimal.NewFromFloat(distance).Mul(rate)

	fmt.Printf("Total fare: $%s\n", fare.StringFixed(2))
}

package main

import (
	"errors"
	"fmt"
	"strings"
)

func getCityFromCode(codeToLookUp string) (City, error) {

	for _, item := range cities {
		if item.code == strings.ToUpper(codeToLookUp) {
			return item, nil
		}
	}

	message := fmt.Sprintf("'%s' is not a valid city code", codeToLookUp)

	return City{"", "", -999, -999}, errors.New(message)
}

func getCabinClassFromCode(codeToLookUp string) (CabinClass, error) {

	for _, item := range cabinClasses {
		if item.code == strings.ToUpper(codeToLookUp) {
			return item, nil
		}
	}

	message := fmt.Sprintf("'%s' is not a valid cabin class code", codeToLookUp)

	return CabinClass{"", "", -999}, errors.New(message)
}

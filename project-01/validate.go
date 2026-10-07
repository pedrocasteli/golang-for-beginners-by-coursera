package main

import (
	"errors"
	"fmt"
)

func getCityFromCode(codeToLookUp string) (City, error) {

	for _, item := range cities {
		if item.code == codeToLookUp {
			return item, nil
		}
	}

	message := fmt.Sprintf("%s is not a valid city code", codeToLookUp)

	return City{"", "", -999, -999}, errors.New(message)
}

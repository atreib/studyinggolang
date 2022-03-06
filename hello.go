package main

import (
	"errors"
	"fmt"
	"log"

	"rsc.io/quote"
)

func GetFullName(name string) (string, error) {
	if name == "" {
		return "", errors.New("name can't be null")
	}

	fullName := name + " Treib"
	return fullName, nil
}

func main() {
	fullName, err := GetFullName("Andre")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello", fullName)
	fmt.Println(quote.Go())
}

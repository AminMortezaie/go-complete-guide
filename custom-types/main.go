package main

import "fmt"

type person struct {
	name string
	age  int
}

type customNumber int

type personData map[string]person

func (number customNumber) pow() customNumber {
	return number * 2
}

func main() {
	var people personData = personData{
		"p1": {
			"amin",
			24,
		},
	}
	fmt.Printf("Print names: %v", people)
}

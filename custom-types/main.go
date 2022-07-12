package main

import "fmt"

type person struct {
	name string
	age  int
}

type customNumber int

type personData map[string]person

func (number customNumber) pow(power int) (result customNumber) {
	result = number

	for i := 0; i < power; i++ {
		result = result * number
	}

	return
}

func main() {
	var people personData = personData{
		"p1": {
			"amin",
			24,
		},
	}
	fmt.Printf("Print names: %v", people)
	fmt.Printf("Power of people:%v", customNumber.pow(2, 69))
}

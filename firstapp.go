package main

import "fmt"

func main() {
	luckyNumber := 17
	// diffrence between declaration in two ways
	// if the right side is INT the left would be INT unless it has decimal
	newNumber := luckyNumber / 3
	var newNumber1 float64 = float64(luckyNumber) / 3

	fmt.Println("Hello World!")
	fmt.Println(luckyNumber)
	fmt.Println(newNumber)
	fmt.Println(newNumber1)
}

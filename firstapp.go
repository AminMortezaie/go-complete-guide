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

	//default declaration would be float64
	float := 1.3423534543

	var defaultFloat float64 = 1.232354353453453453
	var smallFloat float32 = 1.23243543543543534

	fmt.Println(float)
	fmt.Println(defaultFloat)
	fmt.Println(smallFloat)

	// using 1 qoutation only for runes
	// if we print RUNE type without STRING keyword before it will output only a number
	var firstRune rune = '&'
	fmt.Println(firstRune)
	fmt.Println(string(firstRune))

}

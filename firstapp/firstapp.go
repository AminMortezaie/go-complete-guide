package main

import (
	"fmt"

	"github.com/aminmortezaie/firstapp/greeting"
)

func main() {
	luckyNumber := 17
	// diffrence between declaration in two ways
	// if the right side is INT the left would be INT unless it has decimal
	newNumber := luckyNumber / 3
	var newNumber1 float64 = float64(luckyNumber) / 3

	fmt.Println(greeting.GreetingText)
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

	// byte is so similar to rune type
	var firstByte byte = 'a'
	fmt.Println(firstByte)

	firstName := "Amin"
	lastName := "Mortezaie"

	//fullName := firstName + " " + lastName

	// another way to concat two different strings
	fullName := fmt.Sprintf("%v %v", firstName, lastName)

	age := 24
	//there are a lot of problems using two different types in terms of concating them
	// to solve these problems we should utilze Printf instead of Println

	fmt.Printf("My name is %v and I'm %v (Type: %T) years old.", fullName, age, age)

}

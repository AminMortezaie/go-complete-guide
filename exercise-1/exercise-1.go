package main

import "fmt"

func main() {
	firstName := "Amin"
	lastName := "Mortezaie"
	currentYear := 2022

	var birthYear int
	birthYear = 1998
	age := currentYear - birthYear

	nextYear := currentYear + 1

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
	fmt.Println(nextYear)
}

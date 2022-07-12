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
	myVar := person{"ali", 18}
	myVar.birthday()
	myVar.birthday()
	myVar.birthday()
	myVar.birthday()
	myVar.birthday()
	fmt.Printf("myVar: %v\n", myVar)
}

// if you want change the value of the variable you need to declare the pointer instead of value
// because Go would copy by value not copy by reference!

func (p *person) birthday() {
	p.age++
}

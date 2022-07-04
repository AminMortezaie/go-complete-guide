package main

import "fmt"

func main() {
	age := 24

	// myAge := &age
	var myAge *int = &age
	fmt.Println(age)
	fmt.Println(myAge)
	fmt.Println(*myAge)

	*myAge = 34
	fmt.Println(age)
	fmt.Println(double(myAge))
	fmt.Println(age)

}

func double(number *int) (result int) {
	result = *number * 2
	*number = 100
	return
}

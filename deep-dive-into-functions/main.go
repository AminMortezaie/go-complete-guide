package main

import "fmt"

var numbers = []int{1, 2, 3}

func main() {

	a := new([4]int)
	fmt.Printf("a: %v\n", &a)

	doubled := transformNumbers(&numbers, triple)
	fmt.Printf("doubled: %v\n", doubled)
	getTransformerFunction(numbers)
	fmt.Printf("numbers: %v\n", numbers)

	doubled1 := createTransformer(2)
	fmt.Printf("doubled1: %v\n", doubled1(33))

}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

// func double(num int) int {
// 	return num * 2
// }

func triple(num int) int {
	return num * 3
}

func getTransformerFunction(num []int) {
	num[0] = 222
	fmt.Printf("numbers: %v\n", numbers)
}

// transformer function.

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

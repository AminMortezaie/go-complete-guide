package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	doubled := transformNumbers(&numbers, triple)
	fmt.Printf("doubled: %v\n", doubled)

}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}

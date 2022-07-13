package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	doubled := doubledNumbers(&numbers)
	fmt.Printf("doubled: %v\n", doubled)

}

func doubledNumbers(numbers *[]int) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, value*2)
	}
	return dNumbers
}

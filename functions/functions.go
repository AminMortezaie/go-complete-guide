package main

import (
	"fmt"
	"math/rand"
)

func main() {

	a, b := generateRandomNumbers()

	sum := add(a, b)
	printNumber(sum)

}

// version 0

// func add(num1 int, num2 int) int {
// 	return num1 + num2
// }

// version 1
// another declartion of functions by creating name for each output and swiftly we can use return alonely
// pay attention to this example
// it will return sum without refering sum after Return keyword
func add(num1 int, num2 int) (sum int) {
	sum = num1 + num2
	return
}

func printNumber(number int) {
	fmt.Printf("The number is %v", number)
}

func generateRandomNumbers() (r1 int, r2 int) {
	r1 = rand.Intn(10)
	r2 = rand.Intn(6)
	return
}

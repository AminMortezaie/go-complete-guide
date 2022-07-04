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

// func add(num1 int, num2 int) int {
// 	return num1 + num2
// }

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

func generateRandomNumbers() (int, int) {
	randomNumber1 := rand.Intn(10)
	randomNumber2 := rand.Intn(6)
	return randomNumber1, randomNumber2
}

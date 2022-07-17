package variadic

import "fmt"

func main() {

	arr := []int{1, 2, 4}

	fmt.Printf("Result: %v", sumup(arr))
	fmt.Printf("Result: %v", variadicSumup(1, 43, 5353, 423))
}

func sumup(numbers []int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}

func variadicSumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}

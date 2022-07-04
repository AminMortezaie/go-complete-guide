package main

import (
	"fmt"

	"github.com/aminmortezaie/bmi/info"
)

func main() {

	info.PrintWelcome()

	weight, height := getUsersMetrics()

	bmi := weight / (height * height)

	fmt.Printf("Your BMI is: %.2f", bmi)
}

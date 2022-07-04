package main

import (
	"github.com/aminmortezaie/bmi/info"
)

func main() {

	info.PrintWelcome()

	weight, height := getUsersMetrics()

	bmi := calculateBMI(weight, height)

	printBMI(bmi)
}

func calculateBMI(weight float64, height float64) float64 {
	return weight / (height * height)
}

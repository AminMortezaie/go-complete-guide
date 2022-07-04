package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

const mainTitle = "BMI Calculator"
const separator = "---------------"
const weightPrompt = "Please enter your weight(kg): "
const heightPrompt = "Please enter your height(m): "

func main() {
	fmt.Println(mainTitle)
	fmt.Println(separator)
	fmt.Print(weightPrompt)

	// underscore means: we don't care about second output.
	weightInput, _ := reader.ReadString('\n')

	fmt.Print(heightPrompt)
	heightInput, _ := reader.ReadString('\n')

	if runtime.GOOS == "windows" {
		weightInput = strings.Replace(weightInput, "\r\n", "", -1)
		heightInput = strings.Replace(heightInput, "\r\n", "", -1)
	} else {
		weightInput = strings.Replace(weightInput, "\n", "", -1)
		heightInput = strings.Replace(heightInput, "\n", "", -1)
	}

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	bmi := weight / (height * height)

	fmt.Printf("Your BMI is: %.2f", bmi)
}

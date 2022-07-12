package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	selectedChoice, err := getUserChoice()
	if err != nil {
		fmt.Printf("Invalid input, exiting!")
		return
	}
	if selectedChoice == "1" {
		calculateSumUpToNumber()
	} else if selectedChoice == "2" {
		calculateFactorial()
	} else if selectedChoice == "3" {
		calculateSumManually()
	} else {
		calculateListSum()
	}
}

func calculateSumUpToNumber() {
	fmt.Printf("Please enter your number: ")
	chosenNumber, err := getInputNumber()

	if err != nil {
		fmt.Printf("Invalid number input")
		return
	}

	sum := 0

	for i := 1; i <= chosenNumber; i++ {
		sum = sum + i
	}
	fmt.Printf("Result: %v", sum)
}
func calculateFactorial() {
	fmt.Printf("Please enter your number to calculate fac: ")
	chosenNumber, err := getInputNumber()

	if err != nil {
		fmt.Printf("Invalid number input")
		return
	}

	fac := new(uint64)
	i := new(uint64)

	*fac = 1

	for *i = 1; *i <= uint64(chosenNumber); *i++ {
		*fac = (*fac) * (*i)
	}
	fmt.Printf("Result: %v", *fac)
}
func calculateSumManually() {
	isEnteringNumbers := true
	sum := 0
	fmt.Printf("Keep entering numbers, the program will stop once you enter any other values.\n")
	for isEnteringNumbers {
		chosenNumber, err := getInputNumber()
		sum += chosenNumber
		isEnteringNumbers = err == nil
	}
	fmt.Printf("Result: %v", sum)
}
func calculateListSum() {
	fmt.Printf("Enter numbers split them using comma not whitespace.\n")
	inputNumberList, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("Invalid input!")
		return
	}

	inputNumberList = strings.ReplaceAll(inputNumberList, "\r\n", "")
	inputNumber := strings.Split(inputNumberList, ",")
	sum := 0
	for index, value := range inputNumber {
		num, err := strconv.ParseInt(value, 0, 64)
		sum += int(num)
		if err != nil {
			fmt.Printf("This value is not acceptable.\n")
			continue
		}
		fmt.Printf("Index: %v, Value: %v \n", index, value)

	}

	fmt.Printf("Result: %v", sum)
}

func getInputNumber() (int, error) {
	inputNumber, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	inputNumber = strings.ReplaceAll(inputNumber, "\r\n", "")

	// why we can declare another err variable in this part when it's declared before on the first line of func?

	chosenNumber, err := strconv.ParseInt(inputNumber, 0, 64)
	if err != nil {
		return 0, err
	}
	return int(chosenNumber), nil
}

func getUserChoice() (string, error) {
	fmt.Println("1) Add up all the numbers of to number X")
	fmt.Println("2) Build the factorial up to number X")
	fmt.Println("3) Sum up manually entered numbers")
	fmt.Println("4) Sum up a list of entered numbers")
	fmt.Printf("Please make your choice: ")

	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	userInput = strings.ReplaceAll(userInput, "\r\n", "")

	if userInput == "1" ||
		userInput == "2" ||
		userInput == "3" ||
		userInput == "4" {
		return userInput, nil
	} else {
		return "", errors.New("INVALID INPUT")
	}

}

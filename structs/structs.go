package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY)")
	created := time.Now()

	fmt.Println(firstName, lastName, birthDate, created)
}

func getUserData(promptText string) (cleanedInput string) {
	fmt.Println(promptText)

	userInput, _ := reader.ReadString('\n')
	cleanedInput = strings.Replace(userInput, "\r\n", "", -1)
	return
}

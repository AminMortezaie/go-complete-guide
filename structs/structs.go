package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

type User struct {
	firstName string
	lastName  string
	birthDate string
	created   time.Time
}

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY)")
	created := time.Now()

	newUser := User{firstName, lastName, birthDate, created}

	fmt.Println(newUser)
}

func getUserData(promptText string) (cleanedInput string) {
	fmt.Print(promptText)

	userInput, _ := reader.ReadString('\n')
	cleanedInput = strings.Replace(userInput, "\r\n", "", -1)
	return
}

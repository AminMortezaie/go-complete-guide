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

// Using pointers in terms of using less memory
// we would return pointer as we need in our program
// it is firmly correlated to our software architecture.

func NewUser(fName string, lName string, bDate string) (newUser *User) {
	created := time.Now()
	newUser = &User{fName, lName, bDate, created}
	return
}

// This is just like a class with referring to a method.
// to reduce our code, it will handle like an OOP program for you.
// This is known as receiver

// This func uses a pointer and increadably when you pass a pointer,
// Go acknowledges and will act with your pointer same as your variable and you can call its attributes

func (user *User) printer() {
	fmt.Printf("%v %v", user.firstName, user.lastName)
}

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY)")

	newUser := NewUser(firstName, lastName, birthDate)

	// fmt.Println(newUser)
	newUser.printer()
}

func getUserData(promptText string) (cleanedInput string) {
	fmt.Print(promptText)

	userInput, _ := reader.ReadString('\n')
	cleanedInput = strings.Replace(userInput, "\r\n", "", -1)
	return
}

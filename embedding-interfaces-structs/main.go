package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Storer interface {
	Store(fileName string)
}

type Prompter interface {
	PromptForInput()
}

type PromptStorer interface {
	Prompter
	Storer
}

type userInputData struct {
	input string
}

type user struct {
	firstName string
	lastName  string
	*userInputData
}

func newUser(first string, last string) *user {
	return &user{
		firstName:     first,
		lastName:      last,
		userInputData: &userInputData{},
	}
}

//this func just initialize the struct
func newUserInputData() *userInputData {
	return &userInputData{}
}

func (usr *userInputData) PromptForInput() {
	fmt.Printf("Please enter your data: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.ReplaceAll(input, "\r\n", "")

	if err != nil {
		fmt.Printf("Fetching failed!")
		return
	}

	usr.input = input

}

func (usr *userInputData) Store(fileName string) {
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Printf("Creating Failed!")
		return
	}

	defer file.Close()
	file.WriteString(usr.input)
}

func main() {
	data := newUserInputData()

	amin := newUser("amin", "mortezaie")

	// Pashmammmmm!!!
	amin.PromptForInput()
	handleUserInput(data)
}

func handleUserInput(container PromptStorer) {
	fmt.Printf("Ready to Store:::\n")
	container.PromptForInput()
	container.Store("user1.txt")
}

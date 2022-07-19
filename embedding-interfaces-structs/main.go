package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Storer interface {
	Store()
}

type Prompter interface {
	PromptForInput()
}

type userInputData struct {
	input string
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

func (usr *userInputData) Storer(fileName string) {
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Printf("Creating Failed!")
		return
	}

	defer file.Close()
	file.WriteString(usr.input)
}

func main() {

}

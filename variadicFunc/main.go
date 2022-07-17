package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	storeData(getUserInput())

}

func getUserInput() (userInput string) {
	fmt.Printf("Please write down your text: ")
	reader := bufio.NewReader(os.Stdin)
	userInput, _ = reader.ReadString('\n')
	userInput = strings.ReplaceAll(userInput, "\r\n", "")
	fmt.Printf("userInput: %v\n", userInput)
	return
}

func storeData(data string) {
	file, err := os.Create("data.txt")

	if err != nil {
		fmt.Printf("Creating file failed!")
		return
	}

	// using defer keyword in terms of calling something in your func body and you want to execute them at the end of executing.
	defer func() {
		file.Close()
		fmt.Printf("file was not closed correctly!")
	}()

	file.WriteString(data)
	fmt.Printf("Data is created successfully!")

}

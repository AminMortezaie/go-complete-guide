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

	// using panic in order to handle an error completely and if has error, program will crash.
	if err != nil {
		fmt.Printf("Creating file failed!")
		panic(err)
	}

	// using defer keyword in terms of calling something in your func body and you want to execute them at the end of executing.
	// declaring func like this is called anonymous func.
	defer func() {
		file.Close()
		fmt.Printf("file was not closed correctly!")
	}()

	file.WriteString(data)
	fmt.Printf("Data is created successfully!")

}

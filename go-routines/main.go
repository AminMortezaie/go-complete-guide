package main

import (
	"fmt"
	"os"
)

func main() {
	great()

	// declaring channels
	channel1 := make(chan int)
	channel2 := make(chan int)

	// pass channals as a parameter
	go storeMoreData(50000, "50000_1.txt", channel1)
	go storeMoreData(50000, "50000_2.txt", channel2)

	//using channels
	<-channel1
	<-channel2
}

func great() {
	fmt.Printf("Hi there!\n")
}

func storeData(storableText string, fileName string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Printf("Creating Failed!\n")
		return
	}
	defer file.Close()

	_, err = file.WriteString(storableText)

	if err != nil {
		fmt.Printf("Writing Failed!")
		return
	}
}

// add channel
func storeMoreData(lines int, fileName string, c chan int) {
	for i := 0; i < lines; i++ {
		text := fmt.Sprintf("Line %v - Dummy Data \n", i)
		storeData(text, fileName)
	}
	fmt.Printf("-- Done Storing %v lines of Data!\n", lines)
	// initilize channel
	c <- 1
}

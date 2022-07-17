package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

type logData struct {
	message  string
	fileName string
}

func (lData *logData) log() {
	err := ioutil.WriteFile(lData.fileName, []byte(lData.message), fs.FileMode(0644))

	if err != nil {
		fmt.Printf("Storing failed!")
	}
}

type loggableString string

func (text loggableString) log() {
	fmt.Println(text)
}

func main() {
	userLog := logData{"User logged in", "user-log.txt"}
	userLog.log()

	message := loggableString("[INFO] User created.")
	message.log()
}

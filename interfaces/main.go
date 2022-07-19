package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

type logger interface {
	log()
}

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
	createLog(&logData{"user logged out", "user-log.txt"})
	outputValue(message)
	// newNumber := []int{1, 2, 3, 4}
	double([...]int{1, 2, 4})
}

func createLog(data logger) {
	data.log()
}

// it accepts any type of value.
func outputValue(value interface{}) {
	val, ok := value.(loggableString)
	fmt.Printf("val: %v\n", val)
	fmt.Printf("ok: %v\n", ok)
	fmt.Printf("value: %v \n", value)
}

func double(value interface{}) interface{} {
	switch val := value.(type) {
	case int:
		return val * 2
	case float32:
		return val * 2
	case []int:
		newNum := append(val, val...)
		fmt.Printf("newNum: %v\n", newNum)
		return newNum
	case [4]int:
		fmt.Println("This is an array!")
		return ""
	default:
		return ""
	}
}

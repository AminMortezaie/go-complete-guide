package main

import "fmt"

func main() {

	number := new(int)
	fmt.Println(number)
	fmt.Println(*number)

	hobbies := make([]string, 0, 10)
	fmt.Printf("hobbies: %v", cap(hobbies))
}

package main

import "fmt"

func main() {

	number := new(int)
	fmt.Println(number)
	fmt.Println(*number)

	// only for map, slice and channels.
	hobbies := make([]string, 0, 10)
	fmt.Printf("hobbies: %v", cap(hobbies))
}

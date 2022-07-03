package main

import "fmt"

func main() {
	var PI float32 = 3.14
	radius := 5
	var circuleCircumference float32 = 2 * PI * float32(radius)
	fmt.Println(circuleCircumference)
	fmt.Printf("For a radius of %v, the circle circumference is %v", radius, circuleCircumference)
}

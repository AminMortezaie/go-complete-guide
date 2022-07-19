package main

import (
	"fmt"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().Unix())
var randN = rand.New(source)

func main() {

	c := make(chan int)

	go generateValue(c)
	go generateValue(c)

	//getting data from channel
	x := <-c
	y := <-c

	sum := x + y
	fmt.Printf("sum: %v \n", sum)
}

func generateValue(c chan int) {
	sleepTime := randN.Intn(3)
	time.Sleep(time.Duration(sleepTime) * time.Second)

	c <- randN.Intn(10)
}

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
	// set a limiter and make channels into buffered channel.
	// the second parameter is capacity of the buffered channel.
	limiter := make(chan int, 3)

	go generateValue(c, limiter)
	go generateValue(c, limiter)
	go generateValue(c, limiter)
	go generateValue(c, limiter)

	//getting data from channel
	// x := <-c
	// y := <-c

	//using loop for goroutines
	sum := 0
	i := 0

	for num := range c {
		sum += num
		i++

		if i == 4 {
			close(c)
		}
	}

	fmt.Printf("sum: %v \n", sum)
}

func generateValue(c chan int, limit chan int) {
	limit <- 1
	fmt.Printf("Generating value... \n")
	// sleepTime := randN.Intn(3)
	time.Sleep(time.Duration(4) * time.Second)

	c <- randN.Intn(10)

	// if we don't mention this line, it wouldn't run anymore and it will give an error.
	// because it wants to get the output to determine that previous channels was completed their work.
	<-limit
}

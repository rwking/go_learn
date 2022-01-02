package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

func hello(done chan string) {
    n := rand.Intn(100)
	if n < 6 {
		done <- strconv.Itoa(n)
	} else {
		fmt.Println(n, " is not less than 6!")
		hello(done)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	done := make(chan string)
	go hello(done)
	n := <-done
	fmt.Println("Found a number that is < 6: ", n)
}
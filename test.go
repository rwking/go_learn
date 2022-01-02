package main

import (
	"fmt"
	// "time"
)

func main() {

	message := make(chan string)

	go func() {
		fmt.Println("Hello, world!")
		message <- "done"
	}()
	<-message

}
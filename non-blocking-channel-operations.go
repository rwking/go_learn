package main

import (
	"fmt"
	"time"
)


func yeah(m chan string) {
	m <- "yeahhh"
}

func main() {
    messages := make(chan string)
    signals := make(chan bool)

	go yeah(messages) 
	
	time.Sleep(time.Second * 2)
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }

    msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }

    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}
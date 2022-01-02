package main

import (
    "fmt"
    "time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func waitF(from string) {
	// time.Sleep(time.Second)
	for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}


func main() {

    f("direct")
	waitF("Sleep")

    go f("goroutine")
	go waitF("Sleeping")
    go func(msg string) {
        fmt.Println(msg)
    }("going")

    time.Sleep(time.Second*3)
    fmt.Println("done")
}
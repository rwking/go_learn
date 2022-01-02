
// The premise here is that by blocking until the goroutine is complete,
// we're ensuring the program doesn't exit before the goroutine executes.
// This is synchronizing the goroutine with main so they can both finish.

package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func main() {

    done := make(chan bool, 1)
    go worker(done)

    <-done
}
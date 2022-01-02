package main

import "fmt"

func intSeq(n int) func() int {
    return func() int {
        n++
        return n
    }
}

func main() {

    nextInt := intSeq(2)

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := intSeq(10)
    fmt.Println(newInts())
	fmt.Println(newInts())
}
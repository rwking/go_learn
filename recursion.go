package main

import "fmt"

func fact(n int) int {
    if n == 0 {
        return 1
    }
	// fmt.Println(n)
    return n * fact(n-1)
}

func main() {
    fmt.Println(fact(7))

    var fib func(n int) int

    fib = func(n int) int {
        if n < 2 {
            return n
        }
		fmt.Println("Pass ", n, "; Get: ", n-1, n-2)
        return fib(n-1) + fib(n-2)
    }

	// fmt.Println("Pass ", n, "; Get: ", n-1, n-2)

	// fmt.Println(fib(2)) // makes sense: 1+0
	// fmt.Println(fib(3)) // makes sense: 1+1
	// fmt.Println(fib(4)) //
	// fmt.Println(fib(5))
	fmt.Println(fib(200))
    // fmt.Println(fib(7))
	// fmt.Println(fib(8))
	// fmt.Println(fib(9))
}
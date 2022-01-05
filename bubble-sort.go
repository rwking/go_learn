package main

import "fmt"

func bubbleSort(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		for j := i+1; j < len(s); j++ {
			// fmt.Println(i, j)
			if s[j] < s[i] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s
}

func main() {

	s := []int{10, 3, 2, 9, 2, 1, 5}

	fmt.Println("Sorting the slice:")
	fmt.Println(s, "\n")

	sorted := bubbleSort(s)

	fmt.Println(("Sorted!"))
	fmt.Println(sorted)
}

package main

import (
	"fmt"
)

func avgFloatSlice(s []float64) (avg float64) {
	
	sum, avg := 0.0, 0.0

		switch len(s) {
		case 0:
			return 0
		default:
			for _, n := range s {
				sum += n 
				avg = sum / float64(len(s))
		}
	}
	return
}

func main() {
	fmt.Println("Let's average a float slice!")

	mySlice := []float64{3.2, 6.4, 9.6}

	fmt.Println(avgFloatSlice(mySlice))


}
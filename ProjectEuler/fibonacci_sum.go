package main

import "fmt"

func main () {
	limit := 4000000
	sum := 0

	// Initialize the first two terms of Fibonacci sequence

	a, b := 1, 2

	for b <= limit {
		// Check if the current term is even 
		if b%2 == 0 {
			sum += b
		}

		// Calculate the next two terms in the Fibonacci sequence

		a, b = b, a + b
	}

	fmt.Println(sum)
}
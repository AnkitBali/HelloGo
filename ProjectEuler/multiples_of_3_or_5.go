package main

import "fmt"

func sumOfMultiples( limit int) int {
	totalSum := 0
	for i := 0; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			totalSum += i
		}
	}
	return totalSum
}

func main() {
	limit := 1000
	result := sumOfMultiples(limit)
	fmt.Println(result)
}
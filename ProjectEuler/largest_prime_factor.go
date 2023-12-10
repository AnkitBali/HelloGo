package main

import "fmt"

func largestPrimeFactor(n int64) int64 {
	var factor int64 = 2

	for n > 1 {
		if n%factor == 0 {
			n /= factor
		} else {
			factor++
		}
	}

	return factor
}

func main() {
	number := int64(600851475143)
	result := largestPrimeFactor(number)
	fmt.Println(result)
}


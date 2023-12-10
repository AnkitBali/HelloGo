package main

import "fmt"

func isPalindrome(n int) bool {
	orignal := n
	reversed := 0

	for n > 0 {
		digit := n % 10
		reversed = reversed*10 + digit
		n /= 10
	}

	return orignal == reversed
}

func largestPalindromeProduct() int {
	largestPalindrome := 0

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			product := i * j
			if isPalindrome(product) && product > largestPalindrome {
				largestPalindrome = product
			}
		}
	}

	return largestPalindrome
}

func main() {
	result := largestPalindromeProduct()
	fmt.Println(result)
}
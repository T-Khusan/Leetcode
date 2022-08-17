package main

import "fmt"

func main() {
	fib := fib(10)
	fmt.Println(fib)
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	a, b, c := 0, 1, 0

	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}

	return c
}

// 0, 1, 2, 3, 4, 5, 6, 7,  8,  9,  10
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55

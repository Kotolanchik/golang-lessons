package main

import "fmt"

func main() {
	fmt.Print(fibonacci(4))
}
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

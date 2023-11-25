package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A)
	fmt.Scan(&B)

	if A > 99 && B > 99 && A > B {
		return
	}

	total_sum := 0

	for i := A; i <= B; i++ {
		total_sum += i
	}

	fmt.Println(total_sum)
}

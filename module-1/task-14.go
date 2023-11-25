package main

import "fmt"

func main() {
	var n, B int
	fmt.Scan(&n)

	total_sum := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&B)

		if B > 9 && B < 100 && B%8 == 0 {
			total_sum += B
		}
	}

	fmt.Println(total_sum)
}

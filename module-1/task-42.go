package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89,
	fibF := 0
	fibS := 1
	fib := fibF + fibS
	flag := false
	for i := 3; fib < n; i++ {
		fibF = fibS
		fibS = fib
		fib = fibF + fibS
		if fib == n {
			fmt.Println(i)
			flag = true
			break
		}
	}
	if !flag {
		fmt.Println(-1)

	}
}

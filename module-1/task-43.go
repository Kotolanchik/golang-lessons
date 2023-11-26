package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var arr = make([]int, 0)
	i := 0
	for n > 0 {
		if n%2 == 0 {
			arr = append(arr, 0)
		} else {
			arr = append(arr, 1)
		}
		n /= 2
		i++
	}
	newDigit := 0
	degree := 1
	for j := 0; j < i-1; j++ {
		degree *= 10
	}

	i--
	for i >= 0 {
		newDigit += arr[i] * degree
		degree /= 10
		i--
	}

	fmt.Print(newDigit)
}

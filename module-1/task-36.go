package main

import "fmt"

func main() {
	var a, size int
	fmt.Scan(&size)

	countZero := 0
	for i := 0; i < size; i++ {
		fmt.Scan(&a)
		if a == 0 {
			countZero++
		}
	}
	fmt.Print(countZero)
}

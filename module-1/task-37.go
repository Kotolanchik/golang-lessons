package main

import "fmt"

func main() {
	var a, size int
	fmt.Scan(&size)

	min := 99999
	countMin := 0
	for i := 0; i < size; i++ {
		fmt.Scan(&a)
		if min > a {
			min = a
			countMin = 1
		} else if min == a {
			countMin++
		}
	}
	fmt.Print(countMin)
}

package main

import "fmt"

func main() {
	var a int
	var size int
	fmt.Scanln(&size)
	var array = make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Scan(&a)
		array[i] = a
	}

	count := 0
	for i := 0; i < size; i++ {
		if array[i] > 0 {
			count++
		}
	}

	fmt.Print(count)
}

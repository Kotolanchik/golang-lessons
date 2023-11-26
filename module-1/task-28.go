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

	for i := 0; i < size; i++ {
		if i%2 == 0 {
			fmt.Print(array[i], " ")
		}
	}
}

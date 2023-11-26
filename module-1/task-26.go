package main

import "fmt"

func main() {

	var size int
	fmt.Scan(&size)
	slice := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Scan(&slice[i])
	}

	fmt.Print(slice[3])
}

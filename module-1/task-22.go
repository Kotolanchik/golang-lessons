package main

import "fmt"

func main() {
	var workArray [10]uint8
	var swapArray [6]uint8
	for i := 0; i < len(workArray); i++ {
		fmt.Scan(&workArray[i])
	}

	for i := 0; i < len(swapArray); i++ {
		fmt.Scan(&swapArray[i])
	}

	for i := 0; i < len(swapArray); i += 2 {
		buff := workArray[swapArray[i]]
		workArray[swapArray[i]] = workArray[swapArray[i+1]]
		workArray[swapArray[i+1]] = buff
	}

	for i := 0; i < len(workArray); i++ {
		fmt.Print(workArray[i], " ")
	}
}

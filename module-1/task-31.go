package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	degree := 100
	newDigit := 0
	for number > 0 {
		newDigit += number % 10 * degree
		number /= 10
		degree /= 10
	}
	fmt.Println(newDigit)
}

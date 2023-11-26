package main

import "fmt"

func main() {
	var digit int
	fmt.Scan(&digit)

	for digit > 9 {
		buffDigit := 0
		for digit > 0 {
			currentDigit := digit % 10
			digit /= 10
			buffDigit += currentDigit
		}
		digit = buffDigit
	}
	fmt.Println(digit)
}

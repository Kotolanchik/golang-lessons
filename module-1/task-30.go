package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	fmt.Println(number/100 + (number/10)%10 + number%10)
}

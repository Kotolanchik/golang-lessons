package main

import "fmt"

func main() {
	var inp int
	fmt.Scan(&inp)

	first := inp / 100
	second := (inp / 10) % 10
	third := inp % 10
	if first != second && first != third && second != third {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

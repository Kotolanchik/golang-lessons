package main

import "fmt"

func main() {
	var inp int
	for fmt.Scan(&inp); inp <= 100; fmt.Scan(&inp) {
		if inp > 9 {
			fmt.Println(inp)
		}
	}
}

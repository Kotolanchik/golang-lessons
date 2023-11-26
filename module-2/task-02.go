package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Print(minimumFromFour(a, b, c, d))

}

func minimumFromFour(variables ...int) int {
	min := 9999
	for _, el := range variables {
		if el < min {
			min = el
		}
	}
	return min
}

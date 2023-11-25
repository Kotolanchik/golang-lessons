package main

import "fmt"

func main() {
	var x, p, y int
	fmt.Scan(&x, &p, &y)

	result := x
	i := 0
	for result <= y {
		result += result * p / 100
		i++
	}
	fmt.Println(i)
}

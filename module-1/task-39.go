package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	max := -9999
	for i := b; i >= a; i-- {
		if i%7 == 0 && i > max {
			max = i
		}
	}

	if max != -9999 {
		fmt.Println(max)
	} else {
		fmt.Println("NO")
	}
}

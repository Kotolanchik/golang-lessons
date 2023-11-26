package main

import "fmt"

func main() {
	var a [3]int = [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [...]int{1, 2, 3}
	d := [3]int{0: 12, 2: 3}

	fmt.Println(a) // [1 2 3]
	fmt.Println(b) // [1 2 3]
	fmt.Println(c) // [1 2 3]
	fmt.Println(d) // [12 0 3]

	e := [3]int{1, 2, 3}
	f := [3]int{1, 2, 3}
	g := [3]int{3, 2, 1}

	fmt.Println(e == f) // true
	fmt.Println(e == g) // false

	var numbers [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println(numbers[0]) // 1
	fmt.Println(numbers[4]) // 5

	numbers[0] = 87

	fmt.Println(numbers[0]) // 87

	n := [5]int{1, 2, 3, 4, 5}
	fmt.Println(n) // [1 2 3 4 5]

	for _, elem := range n {
		elem = 100
		fmt.Println(elem)

		// 100
		// 100
		// 100
		// 100
		// 100
	}
	fmt.Println(n) // [1 2 3 4 5]

	for idx := range n {
		n[idx] = 100
		fmt.Println(n[idx])

		// 100
		// 100
		// 100
		// 100
		// 100
	}
	fmt.Println(n) // [100 100 100 100 100]
}

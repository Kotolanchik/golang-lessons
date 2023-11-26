package main

import "fmt"

func main() {
	a := 200
	b := &a
	*b++
	c := &b
	**c++ // указатель на указатель
	fmt.Println(a)
}
func test(x1 *int, x2 *int) {
	fmt.Print(*x1 * *x2)
}

func test1(x1 *int, x2 *int) {
	t := *x1
	*x1 = *x2
	*x2 = t
	fmt.Print(*x1, " ", *x2)
}

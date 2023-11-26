package main

func main() {

}

func vote(x int, y int, z int) int {
	if x == y {
		return x
	} else if x == z {
		return x
	} else if y == z {
		return y
	}
	return 0
}

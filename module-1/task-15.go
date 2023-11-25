package main

import "fmt"

func main() {
	var inp int
	current_max := -1
	count_max := 0
	for fmt.Scan(&inp); inp != 0; fmt.Scan(&inp) {
		//	fmt.Scan(&inp)

		//if inp == 0 {
		//	break
		//}

		if current_max < inp {
			current_max = inp
			count_max = 1
		} else if current_max == inp {
			count_max++
		}
	}
	fmt.Println(count_max)
}

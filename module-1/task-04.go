package main

import "fmt"

func main() {
	var input_var int
	fmt.Scan(&input_var)
	fmt.Println(fmt.Sprintf("It is %d hours %d minutes.", input_var/30, input_var%30*2))

}

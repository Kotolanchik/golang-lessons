package main

import "fmt"

func main() {
	var sec int
	fmt.Scan(&sec)

	countH := sec / 3600
	countM := (sec - countH*3600) / 60
	//fmt.Println("h ", countH)
	//fmt.Println("m ", countM)
	//fmt.Println("s ", sec-countH*3600-countM*60)
	fmt.Print(fmt.Sprintf("It is %d hours %d minutes.", countH, countM))
}

package main

import (
	"fmt"
)

// В рамках задания
/*func main() {
	var firstNumber, secondNumber int
	fmt.Scanf("%d %d", &firstNumber, &secondNumber)

	buffFn := firstNumber
	degree := 1
	for buffFn > 9 {
		buffFn /= 10
		degree *= 10
	}
	buffSn := secondNumber
	degreeS := 1
	for buffSn > 9 {
		buffSn /= 10
		degreeS *= 10
	}

	for firstNumber > 0 {
		currentFirstDigit := firstNumber / degree
		firstNumber %= degree
		degree /= 10

		buff := secondNumber
		degreeSBuff := degreeS
		for buff > 0 {
			currentSecondDigit := buff / degreeSBuff
			buff %= degreeSBuff
			degreeSBuff /= 10
			if currentSecondDigit == currentFirstDigit {
				fmt.Print(currentFirstDigit)
				fmt.Print(" ")
				break
			}
		}

	}
}*/

// несколько адекватнее
func main() {
	var first_str, second_str string
	fmt.Scanln(&first_str, &second_str)

	for i := 0; i < len(first_str); i++ {
		for j := 0; j < len(second_str); j++ {
			if first_str[i] == second_str[j] {
				fmt.Printf("%c ", first_str[i])
			}
		}
	}

}

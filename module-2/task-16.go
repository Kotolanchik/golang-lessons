package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	var newStr string
	for _, char := range text {
		if countEl(text, char) == 1 {
			newStr += string(char)
		}
	}

	fmt.Print(newStr)
}

func countEl(s string, cur rune) int {
	res := 0
	for _, char := range s {
		if char == cur {
			res++
		}
	}
	return res
}

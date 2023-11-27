package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	if isValidPassword(text) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}

func isValidPassword(password string) bool {
	if len(password) < 5 {
		return false
	}

	for _, char := range password {
		if !(unicode.IsDigit(char) || unicode.Is(unicode.Latin, char)) {
			return false
		}
	}

	return true
}

func isLatinLetter(r rune) bool {
	return ('A' <= r && r <= 'Z') || ('a' <= r && r <= 'z')
}

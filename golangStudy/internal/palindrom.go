package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func userInputArray() []string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите ваши слова")
	scanner.Scan()

	line := scanner.Text()
	words := strings.Fields(line)

	return words
}
func reverseWord(word string) string {
	runes := []rune(word)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func reverseStrings(words []string) []string {
	reversed := make([]string, len(words))

	for i, word := range words {
		reversed[i] = reverseWord(word)
	}

	return reversed
}

func SearchPalindromWords() {

	userWords := userInputArray()
	reversedWords := reverseStrings(userWords)

	for i := range userWords {
		if userWords[i] == reversedWords[i] {
			fmt.Println(userWords[i])
		}
	}

}

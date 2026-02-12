package internal

import (
	"fmt"
)

func UserInputVowel() {

	var temp string
	runes := []rune(temp)

	fmt.Println("Введите вашу строку: ")

	for {
		fmt.Scanln(&temp)
		if temp == "exit" {
			break
		}
		runes = append(runes, []rune(temp)...)
	}

	vowels := searchForVowel(runes)
	fmt.Println("Количество гласных: ", vowels)
}

func searchForVowel(r []rune) int {

	var countVowel int
	vowel := "аеёиоуыэюяaeiou"
	s := []rune(vowel)

	for i := range r {
		for j := range s {
			if r[i] == s[j] {
				countVowel++
			}
		}
	}
	return countVowel

}

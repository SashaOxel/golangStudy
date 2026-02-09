package internal

import (
	"fmt"
)

func UserChoice() {

	var choice int

	fmt.Println("=== Меню: ====")
	fmt.Println("1. Задача на поиск четных чисел в промежутке")
	fmt.Println("2. Задача на подсчет факториала числа")
	fmt.Println("3. Проверка палиндрома")
	fmt.Println("4. Подсчет гласных")

	fmt.Println("Выберите действие (1-4):")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		UserInputEven()
	case 2:
		UserFactorial()
	case 3:
		SearchPalindromWords()
	}

}

package internal

import "fmt"

func UserFactorial() {
	var userNumber = 0
	var temp = 1

	fmt.Println("Введите ваше число для подсчета факториала: ")
	fmt.Scanln(&userNumber)

	for i := 1; i <= userNumber; i++ {
		temp = temp * i
	}
	fmt.Println("Факториал вашего числа: ", temp)
}

package internal

import (
	"fmt"
	"golangstudy/entity"
)

func UserInputEven() {
	var userInputNumber entity.UserNumber

	fmt.Println("Введите нижнюю границу")
	fmt.Scanln(&userInputNumber.MinNumber)

	fmt.Println("Введите верхнюю границу")
	fmt.Scanln(&userInputNumber.MaxNumber)

	fmt.Println("Ваши границы: ", userInputNumber.MinNumber, ",", userInputNumber.MaxNumber)

	SearchForEven(userInputNumber)
}

func SearchForEven(userNum entity.UserNumber) {

	s := make([]int, 0)

	for i := userNum.MinNumber; i <= userNum.MaxNumber; i++ {
		if i%2 == 0 {
			s = append(s, i)
		}
	}
	fmt.Println("Четные числа: ", s)
}

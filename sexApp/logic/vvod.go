package logic

import (
	"fmt"
	"sexApp/entity"
)

var id23 int = 0

func Uservvod() (porson1, porson2 entity.Person) {
	Saveliy := entity.Person{}

	Pasha := entity.Person{}

	fmt.Println("Введите свое имя")
	fmt.Scanln(&Saveliy.Name)

	fmt.Println("Введите свой возраст")
	fmt.Scanln(&Saveliy.Age)

	fmt.Println("Введите свой гендер")
	process()
	fmt.Println("Ваш Id:", id23)
	fmt.Println("")

	fmt.Println("Введите свое имя")
	fmt.Scanln(&Pasha.Name)

	fmt.Println("Введите свой возраст")
	fmt.Scanln(&Pasha.Age)

	fmt.Println("Введите свой гендер")
	process()
	fmt.Println("Ваш Id:", id23)
	return Saveliy, Pasha

}
func process() entity.Gender {
	var pidorasiki string
	fmt.Scanln(&pidorasiki)
	id23 += 1
	return entity.Gender{
		Title: pidorasiki,
		Id:    id23,
	}
}

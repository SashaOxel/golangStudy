package logic

import (
	"fmt"
	"sexApp/entity"
)

func Sex(person entity.Person, person1 entity.Person) {
	if person.Gender.Title != person1.Gender.Title {
		fmt.Println("Смело занимайтесь сексом")
	} else {
		fmt.Println("Застелитесь нахуй")
	}
}

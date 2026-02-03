package main

import (
	"sexApp/logic"
)

func main() {
	person1, person2 := logic.Uservvod()
	logic.Sex(person1, person2)
}

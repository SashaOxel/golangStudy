package main

import "fmt"

type User struct {
	Name        string
	Age         int
	PhoneNumber int
}

func (p User) greetings() {
	fmt.Println("Всем привет, меня зовут:", p.Name)
}
func main() {
	user1 := User{Name: "пончик", Age: 18, PhoneNumber: 12313}

	user1.greetings()
}

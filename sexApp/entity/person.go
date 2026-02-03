package entity

type Person struct {
	Name   string
	Age    int
	Gender Gender
}

type Gender struct {
	Id    int
	Title string
}

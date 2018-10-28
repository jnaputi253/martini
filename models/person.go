package models

type Person struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Age       int    `json:"age" binding:"required"`
}

func NewPerson(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

package structs

import "fmt"

type customer struct{
	firstName	string
	lastName	string
	age			int
}

func (c customer) save()  {
	fmt.Println("Has been added : ", c.firstName, c.lastName,c.age)
}

func Demo2()  {
	c:= customer{firstName: "Engin", lastName: "Demiroğ", age: 35}
	c.save()
}
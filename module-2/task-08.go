package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

// Встраиваемы типы
type Android struct {
	// 	Person - анонимное поле
	Person Person
	Model  string
}

func main() {
	a := new(Android)
	a.Person.Talk()

	// a.Talk() - если анонимное поле

	var b = Android{
		Model: "model",
		Person: Person{
			Name: "name",
		},
	}
	b.Person.Talk()

}

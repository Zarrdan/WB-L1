package main

import "fmt"

type Human struct {
	Name    string
	Surname string
}

type Action struct {
	Genus string
	Human // передаю в структуру Action струкртуру Human
}

func main() {
	Newperson := Action{
		Genus: "Man",
		Human: Human{Name: "Ivan", Surname: "Ivanov"},
	}
	fmt.Printf("The new person is:%s\n", Newperson)
}

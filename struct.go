package main

import "fmt"

type Person struct {
	name string
	age  int
	email string
}

func main() {
	//初始化
	person := Person{"Tom", 30, "tom@gmail.com"}
	fmt.Println(person) //输出 {Tom 30 tom@gmail.com}

	var person2 = Person{name:"jan", age: 24, email:"jan@gmail.com"}
	fmt.Println(person2)

	pPerson := &person

	fmt.Println(pPerson) //输出 &{Tom 30 tom@gmail.com}

	pPerson.age = 40
	person.name = "Jerry"
	fmt.Println(person) //输出 {Jerry 40 tom@gmail.com}
}

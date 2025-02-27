package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	// var p1 Person
	// p1.name = "tom"
	// p1.age = 18
	// p1.sex = "男"
	// fmt.Println(p1)

	var p2 = new(Person)
	p2.Name = "jack"
	p2.Age = 19
	p2.Sex = "男"

	fmt.Printf("值：%#v 类型：%T", p2, p2)
}

package persion

import "fmt"

type person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *person {
	return &person{
		Name: name,
		Age:  age,
	}
}

func (p *person) GetName() string {
	return p.Name
}
func (p *person) GetAge() int {
	return p.Age
}
func (p *person) SetName(name string) {
	p.Name = name
}
func (p *person) SetAge(age int) {
	if age < 0 && age > 100 {
		fmt.Println("年龄错误")
		return
	}
	p.Age = age
}

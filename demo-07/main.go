package main

import (
	calc01 "demo-07/calc"
	"fmt"

	decima "github.com/shopspring/decimal"

	goJson "github.com/tidwall/gjson"

	persion "demo-07/model"
)

func main() {
	sum := calc01.Add(1, 2)

	fmt.Println(sum)

	price, err := decima.NewFromString("1234567890.1234567890")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(price)

	value := goJson.Get(`{"name":"张三","age":18,"sex":"男"}`, "name")

	fmt.Println(value)

	fmt.Println("——————————————")

	p := persion.NewPerson("张三", 18)

	fmt.Println(p.GetName())
	fmt.Println(p.GetAge())
}

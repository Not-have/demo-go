package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 结构体指针
	var stu = Student{
		Name: "tom",
		Age:  18,
	}
	fmt.Printf("%#v\n", stu)

	jsonByte, _ := json.Marshal(stu)
	jsonStr := string(jsonByte)
	fmt.Printf("%v", jsonStr)
}

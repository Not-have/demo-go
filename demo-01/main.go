package main

import "fmt"

func main() {
	a := 9 // 函数内部可以使用 := 声明变量

	fmt.Println(a)

	var num = 1

	fmt.Println(num)

	var str string = "张三"
	fmt.Println(str)

	for key, v := range str {
		/**
		* key 是字符在字符串中的索引（字节偏移），它是一个整数类型
		* v 是字符的 Unicode 码点（rune 类型），使用 %c 来打印它的字符形式
		 */
		fmt.Printf("key=%v value=%c\n", key, v)
	}

	// var info = {
	// 	nmae: "里斯",
	// 	age: 18
	// }

	const name = "李双江时间"

	fmt.Println(str+name, num)

	flag := false

	if flag {
		fmt.Println("值: ", flag)
	} else {
		fmt.Println("值: ", false)
	}

	for i := 0; i < a; i++ {
		fmt.Println("便利: ", i)
		if i == 4 {
			break
			// return 的话，后面都不执行了
			// return
		}
	}

	var arr = []string{"jave", "js", "ts", "go"}

	// 使用 range 获取索引和值
	for index, v := range arr {
		fmt.Println("index:", index) // 打印索引
		fmt.Println("value:", v)     // 打印值
	}

	var arr02 [3]int
	var arr03 [3]string

	fmt.Printf("arr02:%T arr03:%T", arr02, arr03)

}

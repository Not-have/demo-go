package main

import "fmt"

type calc func(int, int) int

func relus(x int) int {
	if x > 0 {
		return x
	} else {
		return 0
	}
}

func fun2() int {
	var num int

	defer func() {
		num++
	}()

	return num
}

func main() {
	fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	fmt.Println(fun2())
}

package main

import (
	"fmt"
)

func base() {
	var (
		grade string = "unknown"
		marks int    = 100
	)

	switch marks {
	case 100:
		grade = "A+"
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 70:
		grade = "C"
	case 60:
		grade = "E"
	}

	fmt.Println(grade)
}

func type_switch() {
	// 根据变量的数据类型执行相应的动作
	var x interface{}
	x = "a"

	switch t := x.(type) {
	case nil:
		fmt.Printf("x 的数据类型: %T\n", t)
	case int:
		fmt.Printf("x 的数据类型: %T\n", t)
	case string:
		fmt.Printf("x 的数据类型: %T\n", t)
	case float32, float64:
		fmt.Printf("x 的数据类型: %T\n", t)
	}
}

func main() {
	base()
	type_switch()
}

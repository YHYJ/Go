/* 条件语句 - switch */

package main

import (
	"fmt"
	"runtime"
	"time"
)

func base_switch() {
	// 只要匹配到一个case就不再执行后面的case包括default
	// default用来应对所有case都不匹配的情况
	fmt.Println("Go runs on: ")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Current system is Linux.")
	case "darwin":
		fmt.Println("Current system is macOS.")
	default:
		fmt.Printf("Current system is %v?\n", os)
	}
}

func switch_with_no_condition() {
	// 没有条件的switch == switch true
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
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
	base_switch()
	switch_with_no_condition()
	type_switch()
}

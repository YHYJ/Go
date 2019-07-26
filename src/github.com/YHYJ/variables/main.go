package main

import "fmt"

/* 标准的变量声明方式 */

// 声明变量但不初始化
var A string

// 声明变量并初始化
var x int = 1
var q, t int = 1, 2

var (
	a int     = 1
	b int     = 2
	c float32 = 3.0
)

/* 短变量声明 */
// 短变量声明必须给一个明确类型的值
// 函数外的每个语句必须以关键字(var、func等)开始，所以短变量声明不能用在函数之外
func main() {
	p := 1
	fmt.Println(p)
}

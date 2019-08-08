/* 变量 */
// `var` 语句定义变量
// 在函数中，简洁赋值语句`:=`在类型明确时可以提到`var`进行变量声明

package main

import "fmt"

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

/* 短声明变量 */
// 函数外的每个语句必须以关键字开始(var、func等)，所以`:=`不能用在函数之外
func main() {
	p := 1 // 短声明变量必须给一个明确类型的值
	fmt.Println(p)
}

/* `defer`语句 */
// `defer`语句会延迟其后给的函数的执行直到上层函数return

package main

import (
	"fmt"
)

// defer
func basic_defer() {
	// defer语句后面跟的必须是函数调用
	// defer语句会将函数延迟到defer外部函数返回之后执行
	// 延迟调用的函数的参数会立即生成，但直到外部函数返回前该函数都不会被调用
	defer fmt.Println("world")
	fmt.Println("Hello")
}

// defer 栈
func defer_multi() {
	// 延迟的函数调用会被压入一个栈中
	// 当外层函数返回时，会按照后进先出的顺序调用被延迟的函数
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func main() {
	basic_defer()
	defer_multi()
}

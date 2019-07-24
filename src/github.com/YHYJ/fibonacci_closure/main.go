/* 斐波纳契闭包 */

package main

import (
	"fmt"
)

// fib函数的返回值是一个函数func，func函数返回连续的斐波纳契数
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fib()
	fmt.Println(f(), f(), f(), f(), f(), f())
}

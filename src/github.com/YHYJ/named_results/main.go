/* 具名返回值 */
// 指定返回值的名字，应当具有一定的意义，作为文档使用
// 没有参数的return语句返回具名返回值，不推荐用在长函数中，影响可读性

package main

import (
	"fmt"
)

func split(sum int) (x, y int) {
	x=sum*4/9
	y=sum-x
	// return
	return x, y
}

func main() {
	fmt.Println(split(17))
}

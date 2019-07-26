/* 导出名 */
// 在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的
// 例如，Pi 就是个已导出名，它导出自 math 包
// 而pi 并未以大写字母开头，所以是未导出的
// 任何“未导出”的名字在该包外均无法访问

package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(math.Pi)
	fmt.Println(math.Pi)
}

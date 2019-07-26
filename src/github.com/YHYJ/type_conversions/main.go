// 最好使用strconv包

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z, f)
	fmt.Println(string(6))
	fmt.Println(strconv.Itoa(6))
}

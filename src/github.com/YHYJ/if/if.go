package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // math.Pow(x, y float64)返回x的y次方值(float64)
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println("---------------------------")
	fmt.Println( // 如果两次对pow的调用在一个Println里
		// 那么在Println调用之前，两次对pow的调用就执行完毕并返回结果
		// 第一个pow返回9，但此时Println没有调用所以9暂存；
		// 第二个pow先实时27 >= 20打印然后返回20
		// 所以看到的效果是先输出27 >= 20,在输出9 20
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println("---------------------------")
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))
}

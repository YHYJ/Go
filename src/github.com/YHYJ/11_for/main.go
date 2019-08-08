/* for循环语句 */
// go只有一种循环结构：`for`循环
// go中的for集合了for和while的功能

package main

import (
	"fmt"
	"time"
)

// 基本的`for`循环
func for_1() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

// 前置和后置语句为空的`for`循环
func for_2() int {
	sun := 1
	for sun < 1000 {
		sun += sun
	}
	return sun
}

// 不带循环条件的`for`是一个死循环
func for_3() {
	for {
		fmt.Println("Loop...")
		time.Sleep(time.Duration(3) * time.Second)
		// time.Sleep的参数d类型为Duration(type Duration int64)，time定义了最小时间单位为1纳秒
		// 所以3秒要*1000,000,000才是真正的秒
	}
}

func main() {
	sum := for_1()
	sun := for_2()
	fmt.Printf("sum = %v\nsun = %v\n", sum, sun)
	for_3()
}

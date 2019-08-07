// go中没有while循环，go中的for集合了for和while的功能
// time/time.go中定义了：
// const (
//		Nanosecond Duration = 1
//		Microsecond = 1000 * Nanosecond
//		Millisecond = 1000 * Microsecond
//		Second = 1000 * Millisecond
//		Minute = 60 * Second
//		Hour = 60 * Minute
// )

package main

import (
	"fmt"
	"time"
)

func for_1() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func for_2() int {
	sun := 1
	for sun < 1000 {
		sun += sun
	}
	return sun
}

func for_3() {
	for {
		fmt.Println("Loop...")
		time.Sleep(time.Duration(3) * time.Second)
		// time.Sleep的参数d类型为Duration，time定义了常量Duration默认值为1
		// 所以3秒要传递3给Duration然后还要*1000,000,000才是真正的秒
	}
}

func main() {
	sum := for_1()
	sun := for_2()
	fmt.Printf("sum = %v\nsun = %v\n", sum, sun)
	for_3()
}

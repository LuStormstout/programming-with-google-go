package main

import (
	"fmt"
	"math"
)

// 作业一：编写一个程序，获取一个浮点数的整数部分。
func main() {
	fmt.Println("Enter a floating point number:")
	var x float64
	_, err := fmt.Scan(&x)
	if err != nil {
		fmt.Println("Scan failed, err:", err)
	}
	a := math.Trunc(x)
	fmt.Println("Truncated number:", a)
}

package main

import (
	"fmt"
	"math"
)

// 作业一：编写一个程序，获取一个浮点数的整数部分。
func main() {
	var inputNumber float64

	fmt.Println("Enter a floating point number:")
	_, err := fmt.Scan(&inputNumber)
	if err != nil {
		fmt.Println("Scan failed, err:", err)
	}

	truncatedNumber := math.Trunc(inputNumber)
	fmt.Println("Truncated number:", truncatedNumber)
}

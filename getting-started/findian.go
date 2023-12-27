package main

import (
	"fmt"
	"strings"
)

/*
编写一个程序，提示用户输入一个字符串。程序将在输入的字符串中搜索字符 'i'、'a' 和 'n'。
如果输入的字符串以字符 'i' 开头，以字符 'n' 结尾，并且包含字符 'a'，则程序应打印“Found!”。
否则，程序应打印“Not Found!”。程序不区分大小写，因此字符的大小写不重要。

示例：对于以下示例输入字符串，“ian”、“Ian”、“iuiygaygn”、“I d skd a efju N”，程序应打印“Found!”。
对于以下字符串，“ihhhhhn”、“ina”、“xian”，程序应打印“Not Found!”。
*/
func main() {

	var inputString string

	fmt.Println("Enter a string:")
	_, err := fmt.Scan(&inputString)

	if err != nil {
		fmt.Println("Scan failed, err:", err)
	}

	inputString = strings.ToLower(inputString)
	if strings.HasPrefix(inputString, "i") && strings.Contains(inputString, "a") && strings.HasSuffix(inputString, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

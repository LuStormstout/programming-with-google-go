package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	useUnicode()
	useStrings()
	useStrconv()
}

// ------------------------------ Unicode --------------------------------
// Unicode 包是 Go 语言的内置包，它提供了一些用于处理 Unicode 字符的函数。
// Go 语言的字符串底层使用 UTF-8 编码，因此它可以处理任何格式的文本数据。
// UTF-8 是一种变长的编码方式，它可以使用 1~4 个字节来表示一个符号，根据不同的符号而变化字节长度。
// UTF-8 编码规则如下：
// 	1 字节：0xxxxxxx
// 	2 字节：110xxxxx 10xxxxxx
// 	3 字节：1110xxxx 10xxxxxx 10xxxxxx
// 	4 字节：11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
//
// 	其中，x 表示可用编码空间。
// 	UTF-8 编码的优点是：
// 		1. 可变长编码，可以根据不同的符号来选择不同长度的编码，节省空间。
// 		2. 兼容 ASCII 编码，ASCII 编码中的字符只需要 1 个字节，而 UTF-8 编码中的字符也只需要 1 个字节，因此它可以兼容 ASCII 编码。
// 		3. 适合存储和传输文本数据。
//
// 	UTF-8 编码的缺点是：
// 		1. 不适合随机访问，如果要访问第 n 个字符，需要从头开始逐个读取。
// 		2. 中文字符需要 3 个字节来表示，而英文字符只需要 1 个字节，因此中文字符在 UTF-8 编码下会占用更多的存储空间。
// 		3. UTF-8 编码的数据长度不固定，这意味着计算机需要根据每个字符的编码长度来分配内存。
//
// 	UTF-8 编码的应用场景：
// 		1. 存储和传输文本数据。
// 		2. 适合在网络上传输文本数据，因为网络传输的数据量较大，而 UTF-8 编码可以节省空间。
// 		3. 适合在内存中存储文本数据，因为 UTF-8 编码可以节省空间。
// 		4. 适合在文本编辑器中编辑文本数据，因为 UTF-8 编码可以节省空间。
// 		5. 适合在数据库中存储文本数据，因为 UTF-8 编码可以节省空间。
// 		6. 适合在程序中处理文本数据，因为 UTF-8 编码可以节省空间。
// 		7. 不适合在计算机中随机访问文本数据，因为 UTF-8 编码的数据长度不固定，计算机无法确定第 n 个字符的位置。

// unicode 包提供了一些用于处理 Unicode 字符的函数。
func useUnicode() {
	unicode.IsDigit('1')  // 判断字符是否为数字
	unicode.IsLetter('a') // 判断字符是否为字母
	unicode.IsSpace(' ')  // 判断字符是否为空白符
	unicode.IsLower('a')  // 判断字符是否为小写
	unicode.IsUpper('A')  // 判断字符是否为大写
	unicode.ToLower('A')  // 字符转小写
	unicode.ToUpper('a')  // 字符转大写
	unicode.IsPunct(',')  // 判断字符是否为标点符号

	// ...
}

// ------------------------------ Strings --------------------------------
// strings 包提供了一些用于操作字符串的函数。
// Go 语言的字符串是一个字节切片，可以通过将其内容封装在双引号中来创建字符串。
// Go 语言的字符串是 Unicode 兼容的，并且是 UTF-8 编码的。
// 字符串是一些字节的集合，通常但不一定是人类可读的文本。
// 字符串是用来处理人类可读的文本数据的，例如：姓名、地址、邮件、代码等。
// 字符串的值是不可变的，也就是说，字符串一旦创建，它们的值就不能被改变。
// 字符串是一种值类型，且值不可变，即创建某个文本后你无法再次修改这个文本的内容，而是会返回一个新的字符串值。
// 字符串底层是通过 byte 数组实现的，因此可以和 []byte 类型相互转换。
// 字符串是不可变的字节序列，因此可以进行切片操作，获取子字符串。
// 字符串的长度是指它包含的字节数，而不是它包含的字符数。
// 字符串的长度可以用内置的 len() 函数获取。
// 字符串的索引从 0 开始，最后一个字符的索引是字符串长度减 1

func useStrings() {
	res1 := strings.Compare("a", "b") // 比较两个字符串，如果 a 小于 b 则返回 -1，如果 a 等于 b 则返回 0，如果 a 大于 b 则返回 1
	fmt.Println(res1)                 // -1 是按照字典顺序比较的

	var a = "I love this world!"
	var b = "world"

	res2 := strings.Contains(a, b) // 判断字符串 a 是否包含子字符串 b
	fmt.Println(res2)              // true

	res3 := strings.Count(a, b) // 统计字符串 a 中包含多少个子字符串 b
	fmt.Println(res3)           // 1

	res4 := strings.HasPrefix(a, "I") // 判断字符串 a 是否以子字符串 b 开头
	fmt.Println(res4)                 // true

	res5 := strings.HasSuffix(a, "!") // 判断字符串 a 是否以子字符串 b 结尾
	fmt.Println(res5)                 // true

	res6 := strings.Index(a, "love") // 返回字符串 b 在字符串 a 中第一次出现的位置，如果没有找到子字符串 b，则返回 -1
	fmt.Println(res6)                // 2

	res7 := strings.Join([]string{"I", "love", "this", "world!"}, " ") // 将字符串数组中的所有元素连接起来，中间使用指定的分隔符分隔
	fmt.Println(res7)                                                  // I love this world!

	res8 := strings.Repeat("I love this world!", 2) // 将字符串重复 n 次
	fmt.Println(res8)                               // I love this world!I love this world!

	res9 := strings.Replace("I love this world!", "love", "hate", 1) // 将字符串中的子字符串 old 替换为 new，n 表示替换的次数，小于 0 表示全部替换
	fmt.Println(res9)                                                // I hate this world!

	res10 := strings.Split("I love this world!", " ") // 将字符串按照指定的分隔符分割成字符串数组
	fmt.Println(res10)                                // [I love this world!]

	res11 := strings.ToLower("I LOVE THIS WORLD!") // 将字符串全部转换为小写
	fmt.Println(res11)

	res12 := strings.ToUpper("i love this world!") // 将字符串全部转换为大写
	fmt.Println(res12)                             // I LOVE THIS WORLD!

	res13 := strings.Trim(" I love this world! ", " ") // 将字符串两端的指定字符去掉
	fmt.Println(res13)                                 // I love this world!

	res14 := strings.TrimLeft(" I love this world! ", " ") // 将字符串左边的指定字符去掉
	fmt.Println(res14)                                     // I love this world!

	res15 := strings.TrimRight(" I love this world! ", " ") // 将字符串右边的指定字符去掉
	fmt.Println(res15)                                      // I love this world!

	// ...
}

// ------------------------------ Strconv --------------------------------
// strconv 包提供了字符串和其他类型之间的转换功能。

func useStrconv() {
	res1, _ := strconv.Atoi("123") // 将字符串转换为 int 类型
	fmt.Printf("字符串的 123 转换为 int 的值：%d 类型为：%T\n", res1, res1)

	res2 := strconv.Itoa(123) // 将 int 类型转换为字符串
	fmt.Printf("int 的 123 转换为字符串的值：%s 类型为：%T\n", res2, res2)

	res3, _ := strconv.ParseBool("true") // 将字符串转换为 bool 类型
	fmt.Printf("字符串的 true 转换为 bool 的值：%t 类型为：%T\n", res3, res3)

	res4 := strconv.FormatBool(true) // 将 bool 类型转换为字符串
	fmt.Printf("bool 的 true 转换为字符串的值：%s 类型为：%T\n", res4, res4)

	res5, _ := strconv.ParseFloat("3.14", 64) // 将字符串转换为 float64 类型
	fmt.Printf("字符串的 3.14 转换为 float64 的值：%f 类型为：%T\n", res5, res5)

	res6 := strconv.FormatFloat(3.14, 'f', 2, 64) // 将 float64 类型转换为字符串
	fmt.Printf("float64 的 3.14 转换为字符串的值：%s 类型为：%T\n", res6, res6)

	res7, _ := strconv.ParseInt("123", 10, 64) // 将字符串转换为 int64 类型
	fmt.Printf("字符串的 123 转换为 int64 的值：%d 类型为：%T\n", res7, res7)

	res8 := strconv.FormatInt(123, 10) // 将 int64 类型转换为字符串
	fmt.Printf("int64 的 123 转换为字符串的值：%s 类型为：%T\n", res8, res8)

	// ...
}

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// ------------------------------ if --------------------------------
	var x int
	x = rand.Intn(8)
	if x >= 5 {
		println("x is greater than 5, x = ", x)
	} else {
		println("x is less than or equal to 5, x = ", x)
	}

	// ------------------------------ switch --------------------------------
	//switch <var> {
	//	case <val1>:
	//		...
	//	case <val2>:
	//		...
	//	default:
	//		...
	//}

	// Go 语言的 switch 语句中，每个 case 语句的最后不需要加 break 关键字。
	fmt.Println("--------------------")
	switch x {
	case 1:
		fmt.Println("x = 1")
	case 2:
		fmt.Println("x = 2")
	case 3:
		fmt.Println("x = 3")
	default:
		fmt.Println("No matching case, x = ", x)
	}

	// ------------------------------ for --------------------------------
	//for <init>; <condition>; <update> {
	//	...
	//}
	fmt.Println("--------------------")
	fmt.Println("for <init>; <condition>; <update> { ... }")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//for <condition> {
	//	...
	//}
	fmt.Println("--------------------")
	fmt.Println("for <condition> { ... }")
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//for {
	//	...
	//}
	fmt.Println("--------------------")
	fmt.Println("for { ... }")
	j := 0
	for {
		fmt.Println(j)
		j++
		if j > 10 {
			break
		}
	}

	// ------------------------------ goto --------------------------------

	// ------------------------------ break --------------------------------

	// ------------------------------ continue --------------------------------

	// ------------------------------ return --------------------------------

	// ------------------------------ defer --------------------------------

	// ------------------------------ panic --------------------------------

	// ------------------------------ recover --------------------------------

}

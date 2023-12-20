package main

import "fmt"

func main() {
	pointer()
	scope()
	scope2()

	var y *int
	y = foo()
	fmt.Printf("%d", y)
}

// ------------------------------ 指针 ------------------------------
// 1. 指针是存储另一个变量的内存地址的变量。指针指向值的内存地址，而不是值本身。指针类型前缀为 *，指针指向的值称为目标。指针的零值是 nil。指针的类型为 *T，T 为指针指向的类型。Go 不支持指针运算。Go 语言中的指针操作只需要记住两个符号：&（取地址）和 *（根据地址取值）。

// 2. Go 语言中的值类型（int、float、bool、string、array、struct）都有对应的指针类型，如：*int、*int64、*string 等。值类型包括：基本类型、数组、结构体。引用类型包括：指针、slice 切片、map、管道 channel、interface 等。

// 3. 在大部分情况下，值类型的变量的值存储在栈中，引用类型的变量的值存储在堆中，或者存储在栈中，但是值本身是指针。值类型的变量的值在内存中的地址不会改变，但指针指向的值是可以改变的。引用类型的变量的内存地址不会改变，但指针指向的值是可以改变的。

// 4. 值类型的变量赋值给另一个变量，会拷贝一份新的值，两个变量的值是完全独立的。引用类型的变量赋值给另一个变量，不会拷贝新的值，两个变量的值是同一个值，任何一个变量的值发生改变，另一个变量的值也会改变。

// 5. 值类型的变量作为参数传递给函数时，传递的是一份新的值，函数中对该值的修改，不会影响到原来的值。引用类型的变量作为参数传递给函数时，传递的是该值的指针，函数中对该值的修改，会影响到原来的值。

// 6. map、slice、channel 这三种引用类型的变量，创建之后就可以直接使用，不用像 C 语言那样先声明后初始化，再使用。

// pointer 指针
func pointer() {
	// 指针
	// & 取地址，是内存中的虚拟地址
	// * 取值，是内存中的实际值

	var x = 1 // 声明一个int类型的变量，值为1
	var y int // 声明一个int类型的变量，默认值为0

	var ip *int // 声明一个int类型的指针

	ip = &x // 指针指向x的地址
	y = *ip // y是指针ip的值，即x的值 1

	ptr := new(int) // new函数创建一个指针, 指向一个int类型的内存地址, 值为0
	*ptr = 3        // 指针ptr的值为3,*ptr 就是指针ptr的值, 即3

	fmt.Println(x, y, ip) // 1 1 0xc0000140b0
}

// ------------------------------ 堆和栈 --------------------------------
/*
	栈（Stack）：
	栈是一种后进先出（Last In,First Out，LIFO）的数据结构，用于存储函数调用的信息，包括局部变量、返回地址等。
	操作包括压栈（Push）和弹栈（Pop），压栈是将数据放入栈中，弹栈是将数据从栈中取出。
	栈的大小是由操作系统决定的，通常是固定的，超过限制会导致栈溢出。
	栈上分配的内存由编译器自动管理，不需要手动管理。
	函数调用结束后，函数的局部变量会被销毁，函数的返回地址也会被销毁，函数的返回值会被拷贝到调用函数的栈上。
	如果函数的返回值是指针，那么指针指向的值的生命周期由调用函数的栈管理。

	堆（Heap）：
	堆是一片动态分配内存的区域，其大小不固定，由操作系统决定。
	堆上分配的内存由程序员手动管理，如果程序员不手动释放堆上分配的内存，可能导致内存泄漏。
	如果程序员手动释放了堆上分配的内存，该内存就不能再使用。继续使用已释放的内存可能导致野指针问题。
*/
/*
	在 Go 语言中，堆和栈的控制主要由运行时系统负责。
	栈（Stack）：
	Go 语言的栈是由运行时系统自动管理的，不需要程序员手动干预。
	栈的大小是有限的，一般在几兆字节到几十兆字节之间，由操作系统决定。
	每个 Go 协程都有一个独立的栈，栈的大小可以根据需要动态伸缩。
	当一个函数被调用时，会在栈上分配一段空间，用于存储局部变量、函数参数以及其他相关信息。
	当函数调用结束时，这段栈空间会被释放。

	堆（Heap）：
	Go 语言的堆是用于存储动态分配的内存的区域，例如通过 new 或 make 创建的对象。
	堆的大小不是固定的，而是由操作系统动态管理。
	堆上分配的内存由 Go 的垃圾回收器（garbage collector）负责管理。
	垃圾回收器会自动识别不再被引用的对象，并回收它们占用的内存。
	开发者无需手动释放堆上的内存，这有助于减轻内存泄漏和野指针等问题。
*/

// ------------------------------ 变量作用域 ------------------------------
// 变量的作用域，就是变量的有效范围
// 即变量在什么范围内可以使用，超出范围就不能使用
// 这个范围就是变量的作用域，也叫变量的生命周期，变量的作用域分为全局变量和局部变量
// 全局变量：在函数体外声明的变量，作用域为整个包内都可以使用
// 局部变量：在函数体内声明的变量，作用域为函数体内部
// 函数的参数也是局部变量，作用域为函数体内部

var x = 4 // 全局变量

// scope 作用域
func scope() {
	// 全局变量 x 虽然不在 scope 函数中定义，但是可以在 scope 函数中使用
	fmt.Printf("全局变量 x = %d \n", x)

	// 报错，局部变量 y 不能在 scope2 函数外使用
	//fmt.Println(y)
}

// scope2 作用域
func scope2() {
	var x = 3 // 局部变量
	var y = 4 // 局部变量

	// 全局变量 x 会被局部变量 x 覆盖，所以这里 x = 3
	fmt.Printf("在 scope2 中 x 被重新定义覆盖了全局变量中的 x 所以这里 x = %d \n", x)

	// 局部变量 y 在 scope2 中定义所以只能在 scope2 中使用
	fmt.Printf("局部变量 y 在 scope2 中定义所以只能在 scope2 中使用 y = %d \n", y)
}

// ------------------------------ 垃圾回收 --------------------------------
// Go 语言中的垃圾回收器（garbage collector）是一个独立的并发程序，用于回收堆上不再使用的内存。
// 垃圾回收器会在程序运行过程中不断地扫描堆，寻找可以回收的内存。
// 垃圾回收器会自动识别不再被引用的对象，并回收它们占用的内存。
// 开发者无需手动释放堆上的内存，这有助于减轻内存泄漏和野指针等问题。

// 函数 foo() 创建了一个局部变量 x，并返回了 x 的地址。
// 虽然 foo() 函数执行完毕后，其作用域内的局部变量通常会被销毁，
// 但是因为 x 的地址被返回并在 main() 函数中被 y 指向，所以 x 的值不会被销毁，它的生命周期将延续到 y 不再使用它为止。
// 这是因为 Go 语言的垃圾回收器会跟踪每个值的引用，只有当一个值不再被引用时，它才会被垃圾回收器回收。
// 所以，这里的 x 的地址和值都不会被销毁。
func foo() *int {
	x := 1
	return &x
}

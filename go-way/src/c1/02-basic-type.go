package main

import "fmt"

/*
基本类型

bool
int
float32/64



整数：

int8（-128 -> 127）
int16（-32768 -> 32767）
int32（-2,147,483,648 -> 2,147,483,647）
int64（-9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807）
无符号整数：

uint8（0 -> 255）
uint16（0 -> 65,535）
uint32（0 -> 4,294,967,295）
uint64（0 -> 18,446,744,073,709,551,615）
浮点型（IEEE-754 标准）：

float32（+- 1e-45 -> +- 3.4 * 1e38）
float64（+- 5 1e-324 -> 107 1e308）

*/

func main() {
	// var a int
	// var b int32
	// a = 15
	// b = a + a // 编译错误  不能隐式转换类型,  int -> int32
	// b = b + 5 // 因为 5 是常量，所以可以通过编译, 5 是int
	var n int16 = 34
	var m int32
	// compiler error: cannot use n (type int16) as type int32 in assignment
	//m = n
	m = int32(n)

	fmt.Printf("32 bit int is: %d\n", m)
	fmt.Printf("16 bit int is: %d\n", n)
}

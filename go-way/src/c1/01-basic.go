package main

// import "fmt" // package implementing formatted I/O
// import "os"
import (
	"fmt"
	"os"
	"runtime"
)

func Sum(a, b int) int {
	return a + b
}

// func FunctionName(a typea, b typeb) (t1 type1, t2 type2) {
// 	return var1, var2
// }

const cc1 = "C"

var v int = 5

type T struct{}

func init() { // initialization of package
}

// valueOfTypeB = typeB(valueOfTypeA) // 类型转换
func extractType() {
	a11 := 5
	b11 := int(a)
}

const Pi = 3.1415926
const b10 string = "abc"
const cc = "adc"

const Ln2 = 0.69314718

// const Ln2= 0.693147180559945309417232121458\
//              176568075500134360255254120680009
const Log2E = 1 / Ln2 // this is a precise reciprocal
const Billion = 1e9   // float constant
const hardEight = (1 << 100) >> 97

const beef, two, ca1 = "eat", 2, "veg"
const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6

// const (
// 	Monday, Tuesday, Wednesday = 1, 2, 3
// 	Thursday, Friday, Saturday = 4, 5, 6
// )

const (
	a = iota
	b = iota
	c = iota
)

type Color int

const (
	RED    Color = iota // 0
	ORANGE              // 1
	YELLOW              // 2
	GREEN               // ..
	BLUE
	INDIGO
	VIOLET // 6
)

var a1, b1 *int

var a2 int
var b2 bool
var str2 string

var (
	a3   int
	b3   bool
	str3 string
)

var a4 int = 15
var b4 = false

// var identifier [type] = value
var a5 int = 15
var i5 = 5
var b5 bool = false
var str6 string = "Go says hello to the world!"

var (
	a7       = 15
	b7       = false
	str      = "Go says hello to the world!"
	numShips = 50
	city     string
)

var n int64 = 2

var (
	HOME   = os.Getenv("HOME")
	USER   = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

func Func1() {
	a8 := 1
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
	fmt.Println(a8)
}

/*

# 值类型和引用类型

## 值类型

基本类型都属于值类型

## 引用类型

更复杂的数据通常会需要使用多个字，这些数据一般使用引用类型保存

# 使用 := 赋值操作符

*/

func main() {
	// fmt.Printf("greek")
	// var goos string = runtime.GOOS
	// fmt.Printf("The operating system is: %s\n", goos)
	// path := os.Getenv("PATH")
	// fmt.Printf("Path is: %s\n", path)

	testVar()
}

func testVar() {
	var a string = "abc"
	fmt.Println("hello, world", a)
}

func testVar2() {
	// var a, b, c int
	// a, b, c = 5, 7, "abc"
	// a, b, c := 5, 7, "abc"
}

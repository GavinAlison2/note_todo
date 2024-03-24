package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"
)

func SumProductDiff(i, j int) (int, int, int) {
	return i + j, i * j, i - j
}

func SumProductDiffN(i, j int) (s int, p int, d int) {
	s, p, d = i+j, i*j, i-j
	return
}

func main() {
	// sum, prod, diff := SumProductDiff(3, 4)
	// fmt.Println("Sum:", sum, "| Product:", prod, "| Diff:", diff)
	// sum, prod, diff = SumProductDiffN(3, 4)
	// fmt.Println("Sum:", sum, "| Product:", prod, "| Diff:", diff)

	// fmt.Print("First example with -1: ")
	// ret1, err1 := MySqrt(-1)
	// if err1 != nil {
	// 	fmt.Println("Error! Return values are: ", ret1, err1)
	// } else {
	// 	fmt.Println("It's ok! Return values are: ", ret1, err1)
	// }

	// fmt.Print("Second example with 5: ")
	// //you could also write it like this
	// if ret2, err2 := MySqrt(5); err2 != nil {
	// 	fmt.Println("Error! Return values are: ", ret2, err2)
	// } else {
	// 	fmt.Println("It's ok! Return values are: ", ret2, err2)
	// }
	// // named return variables:
	// fmt.Println(MySqrt2(5))

	// testDefer()
	// testCurcirDefer()
	// testFibonacci()
	// testFunArgs()
	testCalTime()
}
func MySqrt(f float64) (float64, error) {
	//return an error as second parameter if invalid input
	if f < 0 {
		return float64(math.NaN()), errors.New("I won't be able to do a sqrt of negative number!")
	}
	//otherwise use default square root function
	return math.Sqrt(f), nil
}

//name the return variables - by default it will have 'zero-ed' values i.e. numbers are 0, string is empty, etc.
func MySqrt2(f float64) (ret float64, err error) {
	if f < 0 {
		//then you can use those variables in code
		ret = float64(math.NaN())
		err = errors.New("I won't be able to do a sqrt of negative number!")
	} else {
		ret = math.Sqrt(f)
		//err is not assigned, so it gets default value nil
	}
	//automatically return the named return variables ret and err
	return
}

func printrec(i int) {
	if i > 10 {
		return
	}
	printrec(i + 1)
	fmt.Printf("%d ", i)
}
func blank_identifier() {
	var i1 int
	var f1 float32
	i1, _, f1 = threeValue()
	fmt.Printf("the int: %d, the float: %f \n", i1, f1)
}
func threeValue() (int, int, float32) {
	return 5, 6, 7.1
}

func test_min_max() {
	var min, max int
	min, max = MinMax(4, 2)
	fmt.Printf("Minmium is: %d, Maximum is: %d\n", min, max)
}
func MinMax(a int, b int) (min int, max int) {
	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}
	return
}

// this function changes reply:
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func test_multiply() {
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply) // Multiply: 50
}
func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) } //此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}

// 变长参数

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func testMulArg() {
	x := min(1, 3, 2, 0)
	fmt.Printf("The mininum is: %d \n", x)
	slice := []int{4, 6, 2, 1, 3}
	x = min(slice...)
	fmt.Print("The mininum in the slice is: %d", x)
}

// defer 使用

func defer1() {
	fmt.Printf("In function1 at the top \n")
	defer defer2()
	fmt.Printf("In function1 at the bottom \n")
}

func defer2() {
	fmt.Printf("function2: defered until the end o the calling funcion!")
}

func testDefer() {
	defer1()
}

func testCurcirDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

// built-in function
/*
close
len、cap
new、make
copy、append
panic、recover
print、println
complex、real、imag
*/

// 递归函数
func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return res
}
func testFibonacci() {
	result := 0
	for i := 0; i < 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}
func callback(y int, f func(int, int)) {
	f(y, 2)
}
func testFunArgs() {
	callback(1, Add)
}

// closure
func testClosure1() {
	func() {
		sum := 0
		for i := 1; i < 100; i++ {
			sum += 1
		}
	}()
}

func f_literal() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) }
		//此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}
func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func testClosure2() {
	p2 := Add2()
	fmt.Printf("call add2 for 3 gives: %v\n", p2(3))
	twoAddr := Adder(2)
	fmt.Printf("The result is: %v \n", twoAddr(3))
}

func Adder2() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

func testClosure3() {
	var f = Adder2()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300), " - ")
}

func MakeAddSuffixFactory(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func testFuncFactory() {
	addBmp := MakeAddSuffixFactory(".bmp")
	addJpeg := MakeAddSuffixFactory(".jpeg")

	addBmp("file")
	addJpeg("file")
}

func testCalTime() {
	start := time.Now()
	testFuncFactory()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s \n", delta)
}

package main

import (
	"fmt"

	"chapter7/book/pack1"
)

func main() {
	var test1 string
	test1 = pack1.ReturnStr()
	fmt.Printf("ReturnStr from packge1: %s \n", test1)
	fmt.Printf("Integer from package1: %d \n", pack1.Pack1Int)
	fmt.Printf("Integer from package1: %f \n", pack1.Pack1Float)
}

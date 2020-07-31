package main

import (
	"fmt"
	"reflect"
)

func main() {
	test(1, 2, 3)
	test1(1, "test")
	test555 := []float64{1, 2, 3}
	test1(1, "test", test555...)
}

func test(test ...int) {
	fmt.Println(test)
	fmt.Println(reflect.TypeOf(test))
}

func test1(test int, test1 string, test2 ...float64) {
	fmt.Println(test)
	fmt.Println(reflect.TypeOf(test))
	fmt.Println(test1)
	fmt.Println(reflect.TypeOf(test1))
	fmt.Println(test2)
	fmt.Println(reflect.TypeOf(test2))
}

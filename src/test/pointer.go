package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		amount := 6
		fmt.Println(amount)
		fmt.Println(&amount)
		fmt.Println(reflect.TypeOf(&amount))
	*/

	var test int
	var testPointer *int
	testPointer = &test
	fmt.Println(testPointer)
	fmt.Println(reflect.TypeOf(testPointer), "\n")

	var test1 bool
	testPointer1 := &test1
	fmt.Println(testPointer1)
	fmt.Println(reflect.TypeOf(testPointer1))
	fmt.Println(*testPointer1, "\n")

	var test2 int = 5
	testPointer2 := &test2
	fmt.Println(testPointer2)
	fmt.Println(reflect.TypeOf(testPointer2))
	fmt.Println(*testPointer2)

	double(test2)
	fmt.Println(test2)
	doublePointer(&test2)
	fmt.Println(test2)

}

func doublePointer(number *int) {
	*number *= 2
}

func double(number int) {
	number *= 2
}

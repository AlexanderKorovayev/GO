package main

import (
	"fmt"
	"reflect"
)

func main() {
	var mySlice []string
	mySlice = make([]string, 7)
	mySlice[0] = "test"
	fmt.Println(mySlice[0])
	fmt.Println(mySlice)

	test := make([]int, 5)
	test[0] = 1
	fmt.Println(test[0])
	fmt.Println(test)

	for i := 0; i < len(test); i++ {
		fmt.Println(test[i])
	}
	for _, letter := range test {
		fmt.Println(letter)
	}

	test1 := []float64{1, 2, 3}
	fmt.Println(test1)

	mas := [5]int{1, 2, 3, 4, 5}
	slice := mas[1:3]
	fmt.Println(mas)
	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(mas))
	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 7)
	fmt.Println(slice)
	fmt.Println(mas)

	slice = append(slice, 8, 9)
	fmt.Println(slice)
	fmt.Println(mas)
}

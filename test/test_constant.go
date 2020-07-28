package main

import "fmt"

const test1 = 55

// короткое объявление не допускается за пределами функций
var test2 int = 3

func main() {
	const test int = 5
	fmt.Println(test)
	fmt.Println(test1)
	fmt.Println(test2)
	//test1 = 66
}

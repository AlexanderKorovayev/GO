package main

import "fmt"

func main() {
	var test [5]string
	test[2] = "5"
	fmt.Println(test)
	var test1 [3]int = [3]int{1, 2, 3}
	fmt.Println(test1)
	test2 := [6]float32{1, 2, 3, 4, 5, 6}
	fmt.Println(test2)
	fmt.Printf("%#v\n", test2)
	fmt.Printf("test %v\n", test2)
	fmt.Printf("len is %v\n", len(test2))

	for i := 0; i < len(test2); i++ {
		fmt.Println(test2[i])
	}

	for index, value := range test2 {
		fmt.Println(index, value)
	}

	for _, value := range test2 {
		fmt.Println(value)
	}
}

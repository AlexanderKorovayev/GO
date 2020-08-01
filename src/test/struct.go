package main

import "fmt"

func main() {
	var struct1 struct {
		num  int
		word string
		flag bool
	}

	fmt.Println(struct1)
	fmt.Printf("%#v\n", struct1)

	struct1.num = 5
	struct1.word = "test"
	struct1.flag = true
	fmt.Println(struct1)

}

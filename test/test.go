package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	test := time.Now()
	test1 := test.Year()
	fmt.Println(reflect.TypeOf(test))
	fmt.Println(test)
	fmt.Println(test1)
}

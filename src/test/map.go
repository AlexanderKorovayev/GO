package main

import "fmt"

func main() {
	var dict map[string]int
	dict = make(map[string]int)
	fmt.Println(dict)

	dict1 := make(map[string]int)
	fmt.Println(dict1)

	dict["test"] = 5
	fmt.Println(dict)

	dict2 := map[string]int{"a": 4, "b": 5}
	fmt.Println(dict2)
	fmt.Println(dict2["c"])
	test, ok := dict2["c"]
	fmt.Println(test, ok)

	for key, val := range dict2 {
		fmt.Println(key, val)
	}

	delete(dict2, "a")
	fmt.Println(dict2)
	delete(dict2, "c")
	fmt.Println(dict2)
}

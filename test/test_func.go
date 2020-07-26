package main

import "fmt"

//стр126
func main() {
	fmt.Println(21.0 / 67.0)
	fmt.Printf("%0.2f \n", 21.0/67.0)
	test := schet(13, 4)
	test1 := schet(15, 3)
	fmt.Println(test)
	fmt.Println(test1)
}

func schet(width float64, height float64) float64 {
	area := width * height
	fmt.Printf("area is %0.2f\n", area)
	return area
}

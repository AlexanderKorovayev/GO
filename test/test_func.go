package main

import (
	"fmt"
)

//стр136
func main() {
	fmt.Println(21.0 / 67.0)
	fmt.Printf("%0.2f \n", 21.0/67.0)
	test, err := schet(13, 4)
	test1, err := schet(15, 3)
	fmt.Println(test, err)
	fmt.Println(test1, err)
}

func schet(width float64, height float64) (float64, error) {
	area := width * height
	fmt.Printf("area is %0.2f\n", area)
	if area != 10 {
		//err := errors.New("ошибочка")
		err := fmt.Errorf("ошибочка: %0.2f", area)
		//fmt.Println(err.Error())
		//log.Fatal(err)
		return 0, err
	}
	return area, nil
}

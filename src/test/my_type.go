package main

import "fmt"

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}

func main() {
	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Println("Name:", porsche.name)
	fmt.Println("Top speed:", porsche.topSpeed)

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	fmt.Println("Description:", bolts.description)
	fmt.Println("Count:", bolts.count)

	var value car
	value.name = "test"
	var pointer *car = &value
	// для структур в го через точечную нотацию можно обращаться к полям как по указателю так и просто, ставить
	// звёздочку не обязательно
	fmt.Println(pointer.name)
	// если использовать класический способ то он тут не работает
	//fmt.Println(*pointer.name)
	// классический способ надо использовать так, и это кстати логично
	fmt.Println((*pointer).name)

	myCar := car{name: "Corvette", topSpeed: 337}
	fmt.Println(myCar)

	//но можно не перечислять все поля, тогда они заполнятся значениями по умолчанию
	myCar1 := car{name: "Lada"}
	fmt.Println(myCar1)
}

package main

import "fmt"

type Liters float64
type Gallons float64
type MyType string

func main() {
	var carFuel Gallons
	var busFuel Liters
	// по умолчанию если написать carFuel = 10.0, то тут будет тип флоат, поэтому надо явно приводить к типу как ниже
	carFuel = Gallons(10.0)
	busFuel = Liters(240.0)
	fmt.Println(carFuel, busFuel)
	// это разные типы поэтому операции между ними не прокатят
	//fmt.Println(carFuel - busFuel)
	// но можно преобразовать к одному
	fmt.Println(carFuel - Gallons(busFuel))

	value := MyType("a MyType value")
	value.sayHi("alex")
	anotherValue := MyType("another value")
	anotherValue.sayHi("kate")
}

func (m MyType) sayHi(test string) {
	fmt.Println("Hi", m, test)
}

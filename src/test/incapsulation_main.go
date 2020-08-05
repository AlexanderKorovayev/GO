package main

import (
	"fmt"
	"log"
	"test/incapsulation_lib"
)

func main() {
	date := incapsulation_lib.Date{}
	err := date.SetYear(3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Year)
	date.SetMonth(5)
	date.SetDay(27)
	fmt.Println(date.Year)
	fmt.Println(date)

	//так мы получаем подобие наследования
	// добавив Date го находит этот тип и размещает его поля в структуре Event
	type Event struct {
		Title string
		incapsulation_lib.Date
	}

	test555 := Event{}
	test555.SetYear(50)
	fmt.Println(test555)
	fmt.Println(test555.Year())

}

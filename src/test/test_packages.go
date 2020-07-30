package main

// пакеты можно создавать в рабочей директории GO, она находится в домашней папке
// пользователя
import (
	//"greeting"
	"fmt"
	"test/keyboard"
)

// go по умолчанию ищет пакеты в горуте или гопазе, поэтому по умолчанию не получится создать проект
// где-то и в нём пакеты и что бы go сам их находил

func main() {
	//greeting.Hello()
	//greeting.Hi()
	num, err := keyboard.GetFloat()
	fmt.Println(num, err)

}

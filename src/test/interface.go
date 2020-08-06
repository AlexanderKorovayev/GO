package main

import (
	"fmt"
	"test/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playList(player, mixtape)

	player1 := gadget.TapeRecorder{}
	mixtape1 := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playList(player1, mixtape1)
}

// вот так создаются интерфейсы
type myInterface interface {
	methodWithoutParameters()
	methodWithParameter(float64)
	methodWithReturnValue() string
}

//вот так включается поддержка интерфейса
type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called")
}

func (m MyType) MethodWithParameter(f float64) {
	fmt.Println("MethodWithParameter called with", f)
}

func (m MyType) MethodWithReturnValue() string {
	return "Hi from MethodWithReturnValue"
}

func (my MyType) MethodNotInInterface() {
	fmt.Println("MethodNotInInterface called")
}

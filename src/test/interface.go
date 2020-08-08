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

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	//тут уупадёт потому что плейер не умеет писать
	recorder.Record()
}

func TryOut1(player Player) {
	player.Play("Test Track")
	player.Stop()
	// а тут не упадёт потому что мы говорим что плейер это в данном случае тайп рекордер который умеет писать
	recorder := player.(gadget.TapeRecorder)
	recorder.Record()
}

func TryOut2(player Player) {
	player.Play("Test Track")
	player.Stop()
	// а тут не упадёт потому что мы говорим что плейер это в данном случае тайп рекордер который умеет писать
	recorder, ok := player.(gadget.TapeRecorder)
	// проверяем тот ли тип пришёл
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a TapeRecorder")
	}
}

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playList(player, mixtape)

	player1 := gadget.TapeRecorder{}
	mixtape1 := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playList(player1, mixtape1)

	// упадёт
	TryOut(gadget.TapeRecorder{})
	// не упадёт
	TryOut1(gadget.TapeRecorder{})
	// упадёт так как не внутри приводит не к тому типу
	TryOut1(gadget.TapePlayer{})
	// а тут уже добавлена проверка на тип
	TryOut2(gadget.TapePlayer{})
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

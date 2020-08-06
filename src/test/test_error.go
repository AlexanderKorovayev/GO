package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenFile(fileName string) (*os.File, error) {

	fmt.Println("Opening", fileName)

	return os.Open(fileName)
}

func CloseFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}

func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	// достаточно сделать один вызов что бы получить этот вызов после любого ретёрна
	defer CloseFile(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}

func test_recover() {
	recover()
}

func freakOut() {
	defer test_recover()
	panic("oh no")
	fmt.Println("I won't be run!")
}

//стр 405
func main() {
	numbers, err := GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}

	//defer надо размещать перед рантаймом
	//defer test_recover()

	// так можно вызвать ран тайм ошибку
	//panic("test123")

	// а эта функция ловит рантайм  по принципу try except, но лучше помещать её в defer в тех случаях когда паник и
	// рековер в одной функуии
	//recover()
	freakOut()
	fmt.Printf("Sum: %0.2f\n", sum)
}

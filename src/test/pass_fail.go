// так делаются комментарии
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("enter percent")
	test := 5
	test, test1 := 8, 9
	fmt.Println(test, test1)
	for x := 1; x < 5; x++ {
		fmt.Println(test, test1)
	}
	x := 3
	for x >= 1 {
		fmt.Println(x)
		x--
	}
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	clear_input := strings.TrimSpace(input)
	float_clear_input, err := strconv.ParseFloat(clear_input, 64)
	if err != nil {
		log.Fatal(err)
	}

	if float_clear_input > 60 {
		fmt.Println("хуясе проценты бля")
	} else {
		fmt.Println("лох ебаный нихуя ты не здал")
	}

}

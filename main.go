package main

import (
	"fmt"
	"strings"
)

/*func calculator(a float32, b string, h float32) {
	var answer float32
	switch b {
	case "+":
		answer = a + h
	case "-":
		answer = a - h
	case "*":
		answer = a * h
	case ":":
		if h != 0 {
			answer = a / h
		} else {
			fmt.Println("На ноль делить нельзя!")
			os.Exit(0)
		}
	case "^":
		answer = float32(math.Pow(float64(a), float64(h)))
	}
	fmt.Println(answer)
}*/

func main() {
	/*var a float32
	var b string
	var h float32
	fmt.Println("Введите первое число")
	fmt.Scan(&a)
	fmt.Println("Введите необходимую операцию: +, -, *, :, ^")
	fmt.Scan(&b)
	if b != "+" || b != "-" || b != "*" || b != ":" || b != "^" {
		fmt.Println("Правильно введите необходимое действие!")
		os.Exit(0)
	}
	fmt.Println("Введите второе число")
	fmt.Scan(&h)
	calculator(a, b, h)*/
	var os float32
	var poc float32
	fmt.Println("Введи основание степени")
	fmt.Scan(&os)
	fmt.Println("Введи показатель степени")
	fmt.Scan(&poc)
	poc1 := fmt.Sprint(poc)
	_, after, _ := strings.Cut(poc1, ".")
	bytes := []byte(after)
	for _, symbol := range bytes {
		if string(symbol) != "0" {
			fmt.Println("Показатель степени не целое число")
			return
		}
	}
	poc2 := int(poc)
	os1 := os
	for i := 1; i < poc2; i++ {
		os = os * os1
	}
	fmt.Println("Ответ:", os)
}

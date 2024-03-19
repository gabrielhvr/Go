package main

import (
	"fmt"
	"projeto-teste/calc"
)

func main() {
	//var num
	var num1 int = 6
	num2 := 0

	result, err := calc.Dividir2(num1, num2)

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println(result)
	fmt.Println("hello world!")

}

func dividir(a, b int) int {
	return a / b
}

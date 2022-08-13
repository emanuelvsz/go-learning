package main

import "fmt"

func main() {
	var num1 int = 10
	num2 := 5

	result := soma(num1, num2)
	fmt.Println(result)

	result2 := subtracao(num1, num2)
	fmt.Println(result2)

	result3 := divisao(num1, num2)
	fmt.Println(result3)

	result4 := multiplicacao(num1, num2)
	fmt.Println(result4)
}

func soma(num1, num2 int) int {
	return num1 + num2
}
func subtracao(num1, num2 int) int {
	return num1 - num2
}
func divisao(num1, num2 int) int {
	return num1 / num2
}
func multiplicacao(num1, num2 int) int {
	return num1 * num2
}
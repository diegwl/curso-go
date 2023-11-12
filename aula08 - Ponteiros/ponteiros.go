package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)

	variavel1 = 20
	fmt.Println(variavel1, variavel2)

	// Ponteiro é uma referência de memória
	var variavel3 int = 10
	var ponteiro *int = &variavel3

	// Print do endereço de memória
	fmt.Println(ponteiro)
	// Print do valor
	fmt.Println(variavel3, *ponteiro)

	variavel3 = 100
	fmt.Println(ponteiro)
	fmt.Println(variavel3, *ponteiro)
}
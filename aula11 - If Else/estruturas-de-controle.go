package main

import "fmt"

func main() {
	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else if numero == 15 {
		fmt.Println("Igual a 15")
	} else {
		fmt.Println("Menor que 15")
	}
	
	// if init
	// outroNumero se torna uma variavel local para o escopo do if else
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else {
		fmt.Println("Número é menor que zero")
	}
}
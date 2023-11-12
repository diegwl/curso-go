package main

import "fmt"

func main() {
	var variavel string = "Variável 1"
	variavel2 := "Variável 2"
	
	fmt.Println(variavel)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	const constante1 = "Constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	println(variavel5, variavel6)
}
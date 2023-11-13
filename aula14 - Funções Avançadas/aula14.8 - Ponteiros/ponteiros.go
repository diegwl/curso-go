package main

import "fmt"

// ponteiro como parâmetro
func inverterSinal(numero *int) {
	// fazendo desreferenciação para realizar a conta
	*numero *= -1
}

func main() {
	numero := 20
	// & para passar o endereço de memória
	inverterSinal(&numero)
	fmt.Println(numero)
}
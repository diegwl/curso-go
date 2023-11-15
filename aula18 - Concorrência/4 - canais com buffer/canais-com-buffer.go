package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "Olá Mundo"
	canal <- "Programando em Go!"
	// canal <- "Terceiro Valor"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem, mensagem2)
}
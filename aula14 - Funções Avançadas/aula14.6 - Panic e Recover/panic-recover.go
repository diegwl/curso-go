package main

import "fmt"

func recuperarExecução() {
	if r := recover(); r != nil {
		fmt.Println("Recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecução()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6!")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução")

	fmt.Println(alunoEstaAprovado(8, 6))
	fmt.Println("Pós execução 2")
}
package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main() {
	// não é exatamente uma herança, ele cria um struct dentro de outro,
	// porém eu posso acessar as variaveis diretamente 
	p1 := pessoa{"Diego", "Castan Lopes", 19, 190}
	fmt.Println(p1)

	e1 := estudante{p1, "ADS", "IFSP"}
	fmt.Println(e1)
	fmt.Println(e1.nome, e1.sobrenome)
}
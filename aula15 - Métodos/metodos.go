package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Println(u.nome, "Dentro do metódo salvar")
}

func (u *usuario) aniversario() {
	u.idade += 1
	fmt.Printf("Parabéns %s, você fez %d anos\n", u.nome, u.idade)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) mudarNome(nome string) {
	u.nome = nome
}

func main() {
	usuario1 := usuario{"Diego", 19}
	fmt.Println(usuario1)

	usuario1.salvar()

	usuario1.aniversario()
	fmt.Println(usuario1.idade)

	fmt.Println(usuario1.maiorDeIdade())

	usuario1.mudarNome("Diegwl")
	fmt.Println(usuario1.nome)
}
package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint8
}

func main() {
	fmt.Println("Structs")
	var u usuario
	u.nome = "Diego"
	u.idade = 19

	var end endereco
	end.logradouro = "avenida exemplo"
	end.numero = 10

	u.endereco = end

	fmt.Println(u)

	u2 := usuario{"Diegwl", 19, endereco{"rua exemplo", 10}}
	fmt.Println(u2)

	u3 := usuario{nome:"Diego"}
	fmt.Println(u3)

	u4 := usuario{idade:19}
	fmt.Println(u4)
}
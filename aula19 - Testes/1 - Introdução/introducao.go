package main

import (
	"introducao-a-testes/enderecos"
	"fmt"
)

func main() {
	TipoEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(TipoEndereco)
}
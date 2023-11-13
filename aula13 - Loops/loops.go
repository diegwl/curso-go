package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	// for como while
	for i < 10 {
		i++
		fmt.Printf("Incrementando i: %d\n", i)
		time.Sleep(time.Second)
	}

	// for original
	for j := 0; j < 10; j++ {
		fmt.Printf("Incrementando j: %d\n", j)
		time.Sleep(time.Second)
	}

	// for com range

	nomes := [3]string{"Diego", "Davi", "Aghata"}

	for i, valor := range nomes {
		fmt.Println(i, valor)
	}

	for _, valor := range nomes {
		fmt.Println(valor)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string {
		"nome": "Diego",
		"sobrenome": "Castan",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// for infinito
	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}
}
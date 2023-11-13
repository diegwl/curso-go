package main

import (
	"busca-ip/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}
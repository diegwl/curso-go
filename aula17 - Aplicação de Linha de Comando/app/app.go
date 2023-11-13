package app

import "github.com/urfave/cli"

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca de IPs e Nomes de Servidor na Internet"

	return app
}
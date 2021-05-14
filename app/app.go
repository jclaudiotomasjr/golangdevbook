package app

import (
	"github.com/urfave/cli"
)

//Gerar vai retonar apliação de linha de comando pronta ser exec
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca de Ips e Nome de servidor na interner"

	return app

}

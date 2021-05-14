package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Gerar vai retonar apliação de linha de comando pronta ser exec
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca de Ips e Nome de servidor na interner"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca Ips de endereções na Internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "gobyexample.com",
				},
			},
			Action: buscarIps,
		},
	}
	return app

}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

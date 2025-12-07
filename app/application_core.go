package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// return
func Cli() *cli.App {
	app := cli.NewApp()
	app.Name = "go_app"
	app.Usage = "Requests servers HTTP and see ip address"
	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Get IPs address",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com",
				},
			},
			Action: func(c *cli.Context) {
				host := c.String("host")
				ips, erro := net.LookupIP(host)
				if erro != nil {
					log.Fatal(erro)
				}
				for _, ip := range ips {
					fmt.Println("[!] IPs encontrados: ", ip)
				}
			},
		},
		{

			Name:  "server",
			Usage: "Get domains servers",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com",
				},
			},
			Action: func(c *cli.Context) {
				host := c.String("host")
				serverDomain, erro := net.LookupNS(host)
				if erro != nil {
					log.Fatal(erro)
				}
				for _, servidor := range serverDomain {
					fmt.Println(servidor.Host)
				}

			},
		},
	}
	return app
}

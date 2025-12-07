package main

import (
	"aplicacao_terminal/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.Cli()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}

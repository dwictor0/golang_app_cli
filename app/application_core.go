package app

import (
	"github.com/urfave/cli"
)

// return
func Cli() *cli.App {
	app := cli.NewApp()
	app.Name = "go_app"
	app.Usage = "Requests servers HTTP and see ip address"

	return app
}

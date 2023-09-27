package main

import (
	"github.com/urfave/cli"
	"h-devops/cmd"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "h-devops"
	app.Usage = "Tools to assist devops using CLI"
	app.Version = "0.4.0"

	app.Commands = []cli.Command{
		{
			Name:    "install-nvm",
			Aliases: []string{"i-nvm"},
			Usage:   "Install Node version manager (NVM)",
			Action:  cmd.InstallNVM,
		},
		{
			Name:    "add-sudoers",
			Aliases: []string{"su"},
			Usage:   "Add a user to sudoers file",
			Action:  cmd.AddSudoers,
		},
		{
			Name:    "test",
			Aliases: []string{"test"},
			Usage:   "Function test",
			Action:  cmd.Test,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

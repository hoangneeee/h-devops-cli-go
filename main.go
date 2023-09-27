package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "h-devops"
	app.Usage = "Devops support tool"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name:    "hello",
			Aliases: []string{"h"},
			Usage:   "Prints 'Hello, World!'",
			Action: func(c *cli.Context) error {
				log.Println("Hello, World!")
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

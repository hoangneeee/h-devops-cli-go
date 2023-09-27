package main

import (
	"github.com/urfave/cli/v2"
	"h-devops/cmd"
	"log"
	"os"
	"time"
)

func main() {
	app := &cli.App{
		Name:     "h-devops",
		Version:  "0.4.0",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "VoHoang",
				Email: "levuthanhtung11@gmail.com",
			},
		},
		Usage:   "Tools to assist devops using CLI",
		Suggest: true,
		Commands: []*cli.Command{
			{
				Name:    "nvm",
				Aliases: []string{"nvm"},
				Usage:   "Node version manager (NVM)",
				Subcommands: []*cli.Command{
					{
						Name:    "install",
						Usage:   "Install Node version manager (NVM)",
						Aliases: []string{"i"},
						Action:  cmd.InstallNVM,
					},
				},
			},
			{
				Name:    "add-sudoers",
				Aliases: []string{"su"},
				Usage:   "Add a user to sudoers file",
				Action:  cmd.AddSudoers,
			},
			{
				Name:    "setup-docker-env",
				Aliases: []string{"docker", "d"},
				Usage:   "Setup Docker env",
				Subcommands: []*cli.Command{
					{
						Name:    "install",
						Usage:   "Install Docker and Docker-compose",
						Aliases: []string{"i"},
						Action:  cmd.SetupDockerEnv,
					},
					{
						Name:      "add-user-to-group",
						Usage:     "Add user to group docker",
						Aliases:   []string{"add"},
						ArgsUsage: "<username>",
						Action:    cmd.AddUserToDockerGroup,
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

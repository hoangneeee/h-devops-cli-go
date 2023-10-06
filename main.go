package main

import (
	"github.com/urfave/cli/v2"
	"h-devops/cmd"
	"log"
	"os"
	"time"
)

const VERSION = "0.5.1"
const PhpDefaultVersion = "7.4"

func main() {
	app := &cli.App{
		Name:     "h-devops",
		Version:  VERSION,
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
				Name:    "commands",
				Aliases: []string{"cmd"},
				Usage:   "List of available commands",
				Action:  cmd.ListCommands,
			},
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
				Name:        "postgres-backup-s3",
				Aliases:     []string{"pbs3"},
				Usage:       "Get service backup data postgres to S3",
				Description: "A handy Docker container to periodically backup PostgresSQL to S3",
				Action:      cmd.GetTempPostgresBackupToS3,
			},
			{
				Name:        "elasticsearch-snapshot-s3",
				Aliases:     []string{"ens"},
				Usage:       "Get service elasticsearch snapshot to S3",
				Description: "A handy Docker container to periodically snapshot Elasticsearch to S3",
				Action:      cmd.GetTempEnsToS3,
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
			{
				Name:    "certbot",
				Aliases: []string{"cert", "c"},
				Usage:   "Certbot helper",
				Subcommands: []*cli.Command{
					{
						Name:    "install",
						Usage:   "Install Certbot",
						Aliases: []string{"i"},
						Action:  cmd.InstallCertbot,
					},
					{
						Name:    "auto-renew-guide",
						Usage:   "Auto renew Let's encrypt certificate for Nginx",
						Aliases: []string{"a"},
						Action:  cmd.AutoRenewCertbotGuide,
					},
					{
						Name:    "expiry",
						Usage:   "Check Certificates expiry date",
						Aliases: []string{"ex"},
						Action:  cmd.CertBotCheckExpiry,
					},
				},
			},
			{
				Name:    "php",
				Aliases: []string{"php"},
				Usage:   "PHP helper",
				Subcommands: []*cli.Command{
					{
						Name:    "install",
						Usage:   "Install PHP",
						Aliases: []string{"i"},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "version",
								Aliases: []string{"v"},
								Value:   PhpDefaultVersion, // Default PHP version
								Usage:   "Specify the PHP version to install",
							},
						},
						Action: cmd.InstallPHP,
					},
					{
						Name:    "remove",
						Usage:   "Remove PHP",
						Aliases: []string{"r"},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "version",
								Aliases: []string{"v"},
								Usage:   "Specify the PHP version to remove",
							},
						},
						Action: cmd.RemovePHP,
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

package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"h-devops/helpers"
	"log"
	"os"
	"os/exec"
	"strings"
)

// InstallNVM installs NVM if it is not already installed.
//
// Parameters:
// - c: the *cli.Context object containing command line arguments and options.
//
// Returns:
// - error: an error object if there was an issue installing NVM, otherwise nil.
func InstallNVM(c *cli.Context) error {
	// Check if NVM is already installed
	cmd := exec.Command("nvm", "--version")
	cmd.Env = os.Environ()
	output, err := cmd.CombinedOutput()

	if err == nil && strings.HasPrefix(string(output), "0.") {
		fmt.Println("NVM is already installed.")
		return nil
	}

	// Download and install NVM
	fmt.Println("Installing NVM...")
	cmd = exec.Command("curl", "-o-",
		"https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh")
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	helpers.HandleError(err)

	fmt.Println("NVM installed successfully.")
	fmt.Println("Please restart your shell to start using NVM.")
	return nil
}

// AddSudoers adds a user to the sudoers file.
//
// It takes a *cli.Context parameter and returns an error.
func AddSudoers(c *cli.Context) error {
	username := c.Args().First()
	if username == "" {
		return cli.Exit("Please specify a username to add to sudoers", 1)
	}

	err := helpers.CheckPermissionSudo()
	helpers.HandleError(err)

	// Add the user to sudoers
	sudoersLine := fmt.Sprintf("%s ALL=(ALL:ALL) ALL", username)
	err = helpers.AddToSudoers(sudoersLine)
	helpers.HandleError(err)

	fmt.Printf("User %s added to sudoers.\n", username)
	return nil
}

func SetupDockerEnv(c *cli.Context) error {
	err := helpers.CheckPermissionSudo()
	helpers.HandleError(err)

	fmt.Println("Installing Docker...")
	if err := installDocker(); err != nil {
		log.Fatalf("Error installing Docker: %v\n", err)
	}
	fmt.Println("Docker installed successfully.")

	fmt.Println("Installing Docker Compose...")
	if err := installDockerCompose(); err != nil {
		log.Fatalf("Error installing Docker Compose: %v\n", err)
	}
	fmt.Println("Docker Compose installed successfully.")
	return nil
}

func AddUserToDockerGroup(c *cli.Context) error {
	username := c.Args().First()
	if username == "" {
		return cli.Exit("Please specify a username", 1)
	}

	err := helpers.RunCmd("usermod", "-aG", "docker", username)
	helpers.HandleError(err)

	// Reload group
	err = helpers.RunCmd("newgrp", "docker")
	helpers.HandleError(err)

	return nil
}

// Private function

func installDocker() error {
	err := helpers.RunCmd("sh", "-c", "curl -fsSL https://get.docker.com | sh")
	helpers.HandleError(err)
	return nil
}

func installDockerCompose() error {
	err := helpers.RunCmd("sh", "-c", "curl -L https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose && chmod +x /usr/local/bin/docker-compose")
	helpers.HandleError(err)
	return nil
}

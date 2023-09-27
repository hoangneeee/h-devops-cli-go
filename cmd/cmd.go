package cmd

import (
	"fmt"
	"github.com/urfave/cli"
	"h-devops/helper"
	"os"
	"os/exec"
	"strings"
)

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
	if err != nil {
		return err
	}

	fmt.Println("NVM installed successfully.")
	fmt.Println("Please restart your shell to start using NVM.")
	return nil
}

func AddSudoers(c *cli.Context) error {
	username := c.Args().First()
	if username == "" {
		return cli.NewExitError("Please specify a username to add to sudoers", 1)
	}

	// Check if the program is running with superuser privileges
	if os.Geteuid() != 0 {
		return cli.NewExitError("This program must be run as root (or using sudo).", 1)
	}

	// Check if the user is already in sudoers
	cmd := exec.Command("grep", "-q", fmt.Sprintf("^%s\\s+ALL=(ALL:ALL)\\s+ALL", username), "/etc/sudoers")
	cmd.Env = os.Environ()
	err := cmd.Run()
	if err == nil {
		fmt.Printf("User %s is already in sudoers.\n", username)
		return nil
	}

	// Add the user to sudoers
	sudoersLine := fmt.Sprintf("%s ALL=(ALL:ALL) ALL", username)
	err = helper.AddToSudoers(sudoersLine)
	if err != nil {
		return err
	}

	fmt.Printf("User %s added to sudoers.\n", username)
	return nil
}

func Test(c *cli.Context) error {
	fmt.Println("Testing...")
	return nil
}

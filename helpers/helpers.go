package helpers

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/exec"
)

func AddToSudoers(line string) error {
	// Open sudoers file for appending
	file, err := os.OpenFile("/etc/sudoers", os.O_WRONLY|os.O_APPEND, 0)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Write the new line to sudoers file
	_, err = file.WriteString(line + "\n")
	if err != nil {
		return err
	}

	return nil
}

func CheckPermissionSudo() error {
	if os.Geteuid() != 0 {
		return cli.Exit("This program must be run as root (or using sudo).", 1)
	}
	return nil
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func RunCmd(arg string, args ...string) error {
	cmd := exec.Command(arg, args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Fprintf(os.Stderr, "# ---> %s\n", cmd)
	return cmd.Run()
}

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
			log.Fatal(err)
		}
	}(file)

	// Write the new line to sudoers file
	_, err = file.WriteString(line + "\n")
	if err != nil {
		return err
	}

	return nil
}

// CheckPermissionSudo checks if the program is running as root (or using sudo).
//
// It does this by checking the effective user ID of the current process. If the effective user ID is not 0,
// it returns an error indicating that the program must be run as root (or using sudo).
//
// Returns:
// - error: An error indicating that the program must be run as root (or using sudo).
func CheckPermissionSudo() error {
	if os.Geteuid() != 0 {
		return cli.Exit("This program must be run as root (or using sudo).", 1)
	}
	return nil
}

func CheckCurlExist() error {
	cmd := exec.Command("which", "curl")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return cli.Exit("Please install curl: sudo apt-get install curl", 1)
	}
	return nil
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// RunCmd executes the given command with the provided arguments.
//
// The arg parameter specifies the command to be executed, and the args parameter
// is an optional list of arguments to be passed to the command.
//
// The function returns an error if the command execution fails.
func RunCmd(arg string, args ...string) error {
	cmd := exec.Command(arg, args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Fprintf(os.Stderr, "# ---> %s\n", cmd)
	return cmd.Run()
}

// LoadContentFromFile loads a template from a file.
//
// It takes a filePath string as a parameter and returns a string and an error.
func LoadContentFromFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// WriteConfigToFile writes the given content to a file specified by the filePath.
//
// Parameters:
// - filePath: a string representing the path of the file.
// - content: a string containing the content to be written to the file.
//
// Returns:
// - an error if there was an issue creating or writing to the file.
func WriteConfigToFile(filePath, content string) error {
	file, err := os.Create(filePath)
	HandleError(err)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

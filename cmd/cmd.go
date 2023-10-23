package cmd

import (
	"bytes"
	"fmt"
	"github.com/urfave/cli/v2"
	"h-devops/helpers"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

type configParams struct {
	Enabled  bool
	Port     int
	MaxRetry int
	BanTime  string
}

// ListCommands description of the Go function.
//
// # List of available commands
//
// ListCommands takes a *cli.Context parameter and returns an error.
func ListCommands(c *cli.Context) error {
	helpers.Log("===============================")
	helpers.Log("List of available commands")
	helpers.Log("===============================")
	helpers.Log("Node commands")
	helpers.SubLog("nvm i")
	helpers.Log("Linux commands")
	helpers.SubLog("su <username>")
	helpers.Log("Get template services commands")
	helpers.SubLog("pbs3")
	helpers.SubLog("ens")
	helpers.Log("Docker commands")
	helpers.SubLog("d i")
	helpers.SubLog("d add <username>")
	helpers.Log("Certbot commands")
	helpers.SubLog("cert i")
	helpers.SubLog("cert a")
	helpers.SubLog("cert ex")
	helpers.Log("Fail2Ban commands")
	helpers.SubLog("f2b i")
	helpers.SubLog("f2b c")
	return nil
}

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

	err = helpers.CheckCurlExist()
	helpers.HandleError(err)

	helpers.Log("Installing Docker...")
	if err := installDocker(); err != nil {
		log.Fatalf("Error installing Docker: %v\n", err)
	}
	helpers.Log("Docker installed successfully.")

	helpers.Log("Installing Docker Compose...")
	if err := installDockerCompose(); err != nil {
		log.Fatalf("Error installing Docker Compose: %v\n", err)
	}
	helpers.Log("Docker Compose installed successfully.")
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

func GetTempPostgresBackupToS3(c *cli.Context) error {
	err := helpers.CheckCurlExist()
	helpers.HandleError(err)

	err = helpers.RunCmd("curl", "-o", "docker-compose.yaml", "https://raw.githubusercontent.com/hoangneeee/postgres-backup-s3/master/docker-compose.example.yaml")
	helpers.HandleError(err)
	return nil
}

func GetTempEnsToS3(c *cli.Context) error {
	err := helpers.CheckCurlExist()
	helpers.HandleError(err)

	err = helpers.RunCmd("curl", "-o", "docker-compose.yaml", "https://raw.githubusercontent.com/hoangneeee/elasticsearch-snapshot-s3/master/docker-compose.example.yaml")
	helpers.HandleError(err)
	return nil
}

func InstallCertbot(c *cli.Context) error {
	err := helpers.CheckPermissionSudo()
	helpers.HandleError(err)

	// Check certbot exist
	cmd := exec.Command("which", "certbot")
	_, err = cmd.CombinedOutput()
	if err == nil {
		helpers.Log("Certbot is already installed.")
		return nil
	} else {
		err := helpers.RunCmd("apt-get", "install", "certbot", "python3-certbot-nginx")
		helpers.HandleError(err)
	}

	return nil
}

// AutoRenewCertbotGuide generates a guide on how to set up auto-renewal for Certbot.
//
// Takes a cli.Context as input.
// Returns an error.
func AutoRenewCertbotGuide(c *cli.Context) error {
	helpers.Log("===========")
	helpers.Log("Guide")
	helpers.Log("===========")
	helpers.Log("Typing command:")
	helpers.SubLog("crontab -e")
	helpers.Log("Insert the following line:")
	helpers.SubLog("00 01 01 */3 * certbot renew --post-hook \"systemctl reload nginx\"")
	return nil
}

// CertBotCheckExpiry checks the expiry of the CertBot.
//
// Parameter(s):
// - c: the cli.Context object.
//
// Return type(s):
// - error: any error that occurred during the execution.
func CertBotCheckExpiry(c *cli.Context) error {
	err := helpers.CheckPermissionSudo()
	helpers.HandleError(err)

	// Check certbot exist
	cmd := exec.Command("which", "certbot")
	_, err = cmd.CombinedOutput()
	if err != nil {
		helpers.Log("Please install certbot: sudo h-devops cert i")
		return nil
	}

	err = helpers.RunCmd("certbot", "certificates")
	helpers.HandleError(err)
	return nil
}

// InstallPHP installs a specific version of PHP.
//
// It takes a *cli.Context as a parameter.
// It returns an error.
func InstallPHP(c *cli.Context) error {
	err := helpers.CheckPermissionSudo()
	helpers.HandleError(err)

	phpVersion := c.String("version")
	helpers.Log("Retrieve PHP version " + phpVersion)
	helpers.Log("Installing PHP version " + phpVersion + "...")

	err = helpers.RunCmd("apt-get", "install", "php"+phpVersion)
	helpers.HandleError(err)

	return nil
}

// RemovePHP removes the specified version of PHP.
//
// It takes a *cli.Context as a parameter.
// It returns an error.
func RemovePHP(c *cli.Context) error {
	err := helpers.CheckPermissionSudo()
	helpers.HandleError(err)
	phpVersion := c.String("version")

	if phpVersion == "" {
		return cli.Exit("Please specify a version", 1)
	}

	err = helpers.RunCmd("apt-get", "remove", "php"+phpVersion)
	helpers.HandleError(err)

	err = helpers.RunCmd("apt", "autorclean")
	helpers.HandleError(err)

	err = helpers.RunCmd("apt", "autoremove")
	helpers.HandleError(err)

	return nil
}

// InstallFail2Ban installs the Fail2Ban software using the provided cli.Context.
//
// The function takes a *cli.Context as its parameter and returns an error.
func InstallFail2Ban(c *cli.Context) error {
	err := helpers.CheckPermissionSudo()
	helpers.HandleError(err)

	helpers.Log("Installing Fail2Ban...")
	err = helpers.RunCmd("apt-get", "install", "-y", "fail2ban")
	helpers.HandleError(err)
	helpers.Log("Fail2Ban installed successfully.")
	err = helpers.RunCmd("service", "fail2ban", "start")
	helpers.HandleError(err)
	return nil
}

// ConfigureFail2Ban generates a Fail2Ban config example based on the provided parameters.
//
// Parameters:
// - c: a pointer to a cli.Context object.
//
// Returns:
// - an error if there was a problem generating or writing the Fail2Ban config, or nil if successful.
func ConfigureFail2Ban(c *cli.Context) error {
	err := helpers.CheckPermissionSudo()
	helpers.HandleError(err)

	helpers.Log("Generating Fail2Ban config example...")

	templateContent := `
# Fail2Ban configuration file
# Generated by h-devops-cli

[DEFAULT]
# Ignore IP addresses listed in this file
# ignoreip = 127.0.0.1/8 ::1

# Defines the default backend used to ban/unban IPs.
# banaction = iptables-multiport

# Email configuration (if you want email notifications)
#destemail = your_email@example.com
#sendername = Fail2Ban
#mta = sendmail

[ssh]
enabled = {{.Enabled}}
port = {{.Port}}
filter = sshd
logpath = /var/log/auth.log
maxretry = {{.MaxRetry}}
bantime = {{.BanTime}}

# You can add more jails for other services or applications as needed.
`

	params := configParams{
		Enabled:  true,
		Port:     22,
		MaxRetry: 3,
		BanTime:  "1h",
	}

	tmpl, err := template.New("fail2ban-config").Parse(templateContent)
	helpers.HandleError(err)

	var configContent bytes.Buffer
	err = tmpl.Execute(&configContent, params)
	helpers.HandleError(err)

	err = helpers.WriteConfigToFile("fail2ban.conf", configContent.String())
	helpers.HandleError(err)
	helpers.Log("Fail2Ban config generated successfully.")
	helpers.SubLog("Path: fail2ban.conf")
	helpers.SubLog("Hint: cp fail2ban.conf /etc/fail2ban/local.conf")
	helpers.SubLog("==> sudo systemctl restart fail2ban.service")
	helpers.SubLog("==> sudo fail2ban-client status")
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

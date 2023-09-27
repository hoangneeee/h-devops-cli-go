package helper

import "os"

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

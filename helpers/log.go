package helpers

import "fmt"

func Log(msg string) {
	fmt.Println("[+] " + msg)
}

func Err(msg string) {
	fmt.Println("[x] " + msg)
}

func SubLog(msg string) {
	fmt.Println("   ⠿ " + msg)
}

func SubErr(msg string) {
	fmt.Println("   ⠍ " + msg)
}

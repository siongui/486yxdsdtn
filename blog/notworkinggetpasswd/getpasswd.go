package getpasswd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func TurnOffEcho() error {
	cmd := exec.Command("stty", "-echo")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func TurnOnEcho() error {
	cmd := exec.Command("stty", "echo")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Getpasswd() (pwd string, err error) {
	reader := bufio.NewReader(os.Stdin)
	err = TurnOffEcho()
	if err != nil {
		return
	}
	fmt.Print("Enter Password: ")
	pwd, err = reader.ReadString('\n')
	if err != nil {
		return
	}
	err = TurnOnEcho()
	if err != nil {
		return
	}
	pwd = strings.TrimSpace(pwd)
	return
}

package main

import (
	"bufio"
	"fmt"
	"net/smtp"
	"net/textproto"
	"os"
	"strings"
	"syscall"

	"github.com/jordan-wright/email"
	"golang.org/x/crypto/ssh/terminal"
)

func credentials() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", "", err
	}
	password := string(bytePassword)

	return strings.TrimSpace(username), strings.TrimSpace(password), nil
}

func main() {
	e := &email.Email{
		To:      []string{"someone@example.com"},
		From:    "someone <someone@example.com>",
		Subject: "Test",
		Text:    []byte("Hello World"),
		Headers: textproto.MIMEHeader{},
	}

	username, password, err := credentials()
	if err != nil {
		panic(err)
	}

	fmt.Println("\nsending mail ...")
	if !strings.HasSuffix(username, "@gmail.com") {
		username += "@gmail.com"
	}
	err = e.Send("smtp.gmail.com:587", smtp.PlainAuth("", username, password, "smtp.gmail.com"))
	if err != nil {
		panic(err)
	}
}

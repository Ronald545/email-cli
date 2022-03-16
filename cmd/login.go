package cmd

import (
	"fmt"
	"net/mail"
	"os"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

type emailError struct{}

func (*emailError) Error() string {
	return "invalid email \n"
}

var CmdLogin = &cobra.Command{
	Use:   "login",
	Short: "insert simple auth credentials into stmp",
	RunE: func(cmd *cobra.Command, args []string) error {
		emailAddr := ""
		passwd := ""

		// prompting user to enter email
		fmt.Println("Please enter your email address: ")
		fmt.Scanln(&emailAddr)

		// validating email
		_, err := mail.ParseAddress(emailAddr)
		if err != nil {
			return &emailError{}
		}

		// prompting user to enter password
		fmt.Println("Please enter your password: ")
		bytePass, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return err
		}
		passwd = string(bytePass)

		// storing email and password in file to persist user session
		storagePath, err := os.UserHomeDir()
		storagePath = storagePath + "/.config/email-cli/"
		if _, err := os.Stat(storagePath+".netrc"); os.IsNotExist(err) {
			os.MkdirAll(storagePath, 0700)
			f, err := os.Create(storagePath+".netrc")
			if err != nil {
				return err
			}
			f.Close()
		}

		data := fmt.Sprintf("%v %v", emailAddr, passwd)
		err = os.WriteFile(storagePath+".netrc", []byte(data), 0644)

		if err != nil {
			return err
		}

		return nil
	},
}

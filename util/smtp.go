package util

import (
	"os"
	"strings"
  "fmt"
  "net/smtp"
)

type loginError struct {}
func (*loginError) Error() string {
  return "your login credentials were not found: please run \"email-cli login\" before you send mail"
}

func Send(message string, recipients ...string) error {
  // parse login credentials
  storagePath, err := os.UserHomeDir()
  if err != nil {
    return err
  }
	storagePath = storagePath + "/.config/email-cli/"
	if _, err := os.Stat(storagePath+".netrc"); os.IsNotExist(err) { 
    return &loginError{}
  }
  
  data, err := os.ReadFile(storagePath+".netrc")
  if err != nil {
    return err
  }
  credentials := strings.Fields(string(data))
  
   // smtp server configuration
  smtpHost := "smtp.gmail.com"
  smtpPort := "587"

  // Authentication
  auth := smtp.PlainAuth("", credentials[0], credentials[1], smtpHost)
 
  // Sending email
  err = smtp.SendMail(smtpHost+":"+smtpPort, auth, credentials[0], recipients, []byte(message))
  if err != nil {
    return err
  }

  // prompt user 
  fmt.Println("Email Sent Successfully!")

  return nil
}

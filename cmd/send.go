package cmd

import (
	"net/mail"

	"github.com/Ronald545/email-cli/util"
	"github.com/spf13/cobra"
)

var SendCmd = &cobra.Command{
	Use:   "send",
	Short: "sends email to individual user",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// validate emails that are entered
		_, err := mail.ParseAddress(args[0])

		if err != nil {
			return &emailError{}
		}

		err = util.Send(args[1], args[2], args[0]) // subject, recipients, message
		if err != nil {
			return err
		}
		return nil
	},
}

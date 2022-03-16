package main

import (
	"fmt"
	"github.com/Ronald545/email-cli/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "email cli",
		Short: "email-cli is a fast and flexible email client used to send large amounts of emails with templates and data",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cmd.Short)
		},
	}

	rootCmd.AddCommand(cmd.CmdLogin)
  rootCmd.AddCommand(cmd.SendCmd)
	rootCmd.Execute()
}

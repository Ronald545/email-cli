package cmd 

import (
  "github.com/spf13/cobra"
  "github.com/Ronald545/email-cli/util"
)

var Recipient string
var Message string

var SendCmd = &cobra.Command{
  Use:   "send",
  Short: "sends email to individual user",
  Args: cobra.MinimumNArgs(1),
  RunE: func(cmd *cobra.Command, args []string) error {
    err := util.Send(args[1], args[0])
    if err != nil {
      return err
    }
    return nil
  },
}

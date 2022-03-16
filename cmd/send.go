package cmd 

import (
  "github.com/spf13/cobra"
  "github.com/Ronald545/email-cli/util"
)

var SendCmd = &cobra.Command{
  Use:   "send",
  Short: "sends email to individual user",
  RunE: func(cmd *cobra.Command, args []string) error {
    err := util.Send("Hi There", "spikyron@gmail.com")
    if err != nil {
      return err
    }
    return nil
  },
}


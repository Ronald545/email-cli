package cmd 

import (
  "fmt"

  "github.com/spf13/cobra"
)

var SendCmd = &cobra.Command{
  Use:   "try",
  Short: "Try and possibly fail at something",
  RunE: func(cmd *cobra.Command, args []string) error {
    
  },
}

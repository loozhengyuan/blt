// Package cmd provides the commands for interacting with the application.
package cmd

import (
	"github.com/spf13/cobra"
)

// New returns a new command-line parser.
func New() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:           "bl",
		Short:         "Blocklist management tool.",
		SilenceErrors: true,
		SilenceUsage:  true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}
	return cmd, nil
}

// Package cmd provides the commands for interacting with the application.
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/loozhengyuan/blt/internal/cmd/build"
	"github.com/loozhengyuan/blt/internal/cmd/version"
)

// New returns a new command-line parser.
func New() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:           "blt",
		Short:         "Blocklist management tool.",
		SilenceErrors: true,
		SilenceUsage:  true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}
	cmd.AddCommand(build.New())
	cmd.AddCommand(version.New())
	return cmd, nil
}

// Package build provides the build command.
package build

import (
	"github.com/spf13/cobra"
)

// New returns the build command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Builds blocklist according to a spec file",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return cmd
}

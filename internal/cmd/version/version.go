// Package version provides the version command.
package version

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/loozhengyuan/blt/internal/build"
)

const appName = "blt"

type options struct {
	format string
}

// New returns the version command.
func New() *cobra.Command {
	opts := &options{}
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Prints the current version information",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			info := build.Info(appName)
			switch opts.format {
			case "text":
				if err := info.OutputText(os.Stdout); err != nil {
					return fmt.Errorf("output text: %v", err)
				}
			case "json":
				if err := info.OutputJSON(os.Stdout); err != nil {
					return fmt.Errorf("output json: %v", err)
				}
			default:
				return fmt.Errorf("unknown format value: %s", opts.format)
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&opts.format, "format", "f", "text", "output format")
	return cmd
}

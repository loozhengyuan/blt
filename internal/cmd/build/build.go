// Package build provides the build command.
package build

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/loozhengyuan/blt/internal/blocklist"
	"github.com/loozhengyuan/blt/internal/blocklist/spec"
)

// New returns the build command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Builds blocklist according to a spec file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Read config file
			f, err := os.Open(args[0])
			if err != nil {
				return fmt.Errorf("open file: %w", err)
			}
			defer f.Close()

			// Parse config file
			cfg, err := spec.NewV1SpecFromYAML(f)
			if err != nil {
				return fmt.Errorf("parse cfg: %w", err)
			}

			// Create manifest
			m, err := blocklist.NewManifestFromV1Spec(cfg)
			if err != nil {
				return fmt.Errorf("new manifest from spec: %w", err)
			}

			// Build and export blocklist
			o, err := m.Build()
			if err != nil {
				return fmt.Errorf("build blocklist: %w", err)
			}
			if cfg.Output.Destinations != nil {
				for _, dest := range cfg.Output.Destinations {
					f, err := os.Create(dest.FilePath)
					if err != nil {
						return fmt.Errorf("create export file: %w", err)
					}
					defer f.Close()

					// TODO: Default template
					tmpl, err := template.New("").Parse(dest.CustomTemplate)
					if err != nil {
						return fmt.Errorf("parse tmpl: %w", err)
					}
					if err := tmpl.Execute(f, o); err != nil {
						return fmt.Errorf("execute tmpl: %w", err)
					}
				}
			}
			return nil
		},
	}
	return cmd
}

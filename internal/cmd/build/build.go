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

			// Derive parser from blocklist type
			var p blocklist.Parser
			switch cfg.Kind {
			case "ipbl":
				p, err = blocklist.NewIPParser()
				if err != nil {
					return fmt.Errorf("new ip parser: %w", err)
				}
			case "dnsbl":
				p, err = blocklist.NewFQDNParser()
				if err != nil {
					return fmt.Errorf("new fqdn parser: %w", err)
				}
			default:
				return fmt.Errorf("unknown kind: %v", cfg.Kind)
			}

			// Build blocklist
			bl := blocklist.NewBlocklist()
			if cfg.Policy.Allow.Items != nil {
				bl.Allow(cfg.Policy.Allow.Items...)
			}
			if cfg.Policy.Allow.Includes != nil {
				for _, src := range cfg.Policy.Allow.Includes {
					ref, err := blocklist.NewURLSource(src.URL, p)
					if err != nil {
						return fmt.Errorf("new url ref: %w", err)
					}
					if err := bl.AllowFrom(ref); err != nil {
						return fmt.Errorf("allow from ref: %w", err)
					}
				}
			}
			if cfg.Policy.Deny.Items != nil {
				bl.Deny(cfg.Policy.Deny.Items...)
			}
			if cfg.Policy.Deny.Includes != nil {
				for _, src := range cfg.Policy.Deny.Includes {
					ref, err := blocklist.NewURLSource(src.URL, p)
					if err != nil {
						return fmt.Errorf("new url ref: %w", err)
					}
					if err := bl.DenyFrom(ref); err != nil {
						return fmt.Errorf("deny from ref: %w", err)
					}
				}
			}

			// Export config
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
					if err := tmpl.Execute(f, bl); err != nil {
						return fmt.Errorf("execute tmpl: %w", err)
					}
				}
			}
			return nil
		},
	}
	return cmd
}

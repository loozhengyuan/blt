package blocklist

import (
	"fmt"
	"sort"
	"time"

	"github.com/loozhengyuan/blt/internal/blocklist/spec"
)

// Kind represents the type of a blocklist.
type Kind int

// Enum types for the type of a blocklist.
const (
	KindUnspecified Kind = iota // Default null value
	KindIPBL
	KindDNSBL
)

// Manifest represents the specification for generating a blocklist.
type Manifest struct {
	permitted []Source
	forbidden []Source
}

// Allow registers one or more Source as a permitted source.
func (m *Manifest) Allow(src ...Source) {
	m.permitted = append(m.permitted, src...)
}

// Deny registers one or more Source as a forbidden source.
func (m *Manifest) Deny(src ...Source) {
	m.forbidden = append(m.forbidden, src...)
}

// Build generates and returns the pointer to a new Export.
func (m *Manifest) Build() (*Export, error) {
	// Generate lists
	pl, err := m.permittedList()
	if err != nil {
		return nil, fmt.Errorf("generate permitted list: %w", err)
	}
	fl, err := m.forbiddenList()
	if err != nil {
		return nil, fmt.Errorf("generate forbidden list: %w", err)
	}

	// Build blocklist
	items := make([]string, 0)
	for k := range fl {
		// Append if not excluded by permitted list
		_, ok := pl[k]
		if !ok {
			items = append(items, k)
		}
	}
	sort.Strings(items)
	bl := Export{
		data:      items,
		timestamp: time.Now().UTC(),
	}
	return &bl, nil
}

func (m *Manifest) permittedList() (map[string]int, error) {
	items := make(map[string]int, len(m.permitted))
	for _, src := range m.permitted {
		data, err := src.Items()
		if err != nil {
			return nil, fmt.Errorf("read permitted source: %w", err)
		}
		for _, v := range data {
			items[v] += 1
		}
	}
	return items, nil
}

func (m *Manifest) forbiddenList() (map[string]int, error) {
	items := make(map[string]int, len(m.forbidden))
	for _, src := range m.forbidden {
		data, err := src.Items()
		if err != nil {
			return nil, fmt.Errorf("read forbidden source: %w", err)
		}
		for _, v := range data {
			items[v] += 1
		}
	}
	return items, nil
}

// NewManifest returns the pointer to a new Manifest.
func NewManifest() *Manifest {
	return &Manifest{
		permitted: make([]Source, 0),
		forbidden: make([]Source, 0),
	}
}

// NewManifestFromV1Spec returns a new Manifest from a configuration spec.
func NewManifestFromV1Spec(cfg *spec.V1Spec) (*Manifest, error) {
	m := &Manifest{
		permitted: make([]Source, 0),
		forbidden: make([]Source, 0),
	}

	// Derive parser from blocklist type
	var p Parser
	switch cfg.Kind {
	case "ipbl":
		var err error
		p, err = NewIPParser()
		if err != nil {
			return nil, fmt.Errorf("new ip parser: %w", err)
		}
	case "dnsbl":
		var err error
		p, err = NewFQDNParser()
		if err != nil {
			return nil, fmt.Errorf("new fqdn parser: %w", err)
		}
	default:
		return nil, fmt.Errorf("unknown kind: %v", cfg.Kind)
	}

	// Register permitted sources
	if cfg.Policy.Allow.Items != nil {
		m.Allow(NewListSource(cfg.Policy.Allow.Items...))
	}
	if cfg.Policy.Allow.Includes != nil {
		for _, src := range cfg.Policy.Allow.Includes {
			ref, err := NewURLSource(src.URL, p)
			if err != nil {
				return nil, fmt.Errorf("new url ref: %w", err)
			}
			m.Allow(ref)
		}
	}

	// Register forbidden sources
	if cfg.Policy.Deny.Items != nil {
		m.Deny(NewListSource(cfg.Policy.Deny.Items...))
	}
	if cfg.Policy.Deny.Includes != nil {
		for _, src := range cfg.Policy.Deny.Includes {
			ref, err := NewURLSource(src.URL, p)
			if err != nil {
				return nil, fmt.Errorf("new url ref: %w", err)
			}
			m.Deny(ref)
		}
	}
	return m, nil
}

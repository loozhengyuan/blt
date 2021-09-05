// Package blocklist provides the types to manage a blocklist.
package blocklist

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

// Blocklist represents a list of undesired IPs and FQDNs.
type Blocklist struct {
	mu          sync.RWMutex
	permitted   map[string]int
	forbidden   map[string]int
	lastUpdated time.Time
}

// Allow marks IPs and FQDNs as permitted.
//
// When the blocklist is exported, any IPs or FQDNs in the
// permitted list is omitted.
func (bl *Blocklist) Allow(data ...string) {
	bl.mu.Lock()
	defer bl.mu.Unlock()
	for _, v := range data {
		bl.permitted[v] += 1
	}
	bl.lastUpdated = time.Now().UTC()
}

// AllowFrom is like Allow but reads the set of IPs and FQDNs
// from a Source.
func (bl *Blocklist) AllowFrom(src Source) error {
	items, err := src.Items()
	if err != nil {
		return fmt.Errorf("get items: %w", err)
	}
	bl.Allow(items...)
	return nil
}

// Deny marks IPs and FQDNs as forbidden.
func (bl *Blocklist) Deny(data ...string) {
	bl.mu.Lock()
	defer bl.mu.Unlock()
	for _, v := range data {
		bl.forbidden[v] += 1
	}
	bl.lastUpdated = time.Now().UTC()
}

// DenyFrom is like Deny but reads the set of IPs and FQDNs
// from a Source.
func (bl *Blocklist) DenyFrom(src Source) error {
	items, err := src.Items()
	if err != nil {
		return fmt.Errorf("get items: %w", err)
	}
	bl.Deny(items...)
	return nil
}

// Build generates and returns the pointer to a new Export.
func (bl *Blocklist) Build() *Export {
	bl.mu.RLock()
	defer bl.mu.RUnlock()
	items := make([]string, 0)
	for k := range bl.forbidden {
		_, ok := bl.permitted[k]
		if !ok {
			items = append(items, k)
		}
	}
	sort.Strings(items)
	return &Export{
		data:      items,
		timestamp: time.Now().UTC(),
	}
}

// NewBlocklist returns the pointer to a new Blocklist.
func NewBlocklist() *Blocklist {
	return &Blocklist{
		permitted:   make(map[string]int),
		forbidden:   make(map[string]int),
		lastUpdated: time.Now().UTC(),
	}
}

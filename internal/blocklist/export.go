package blocklist

import (
	"time"
)

// Export represents a generated blocklist.
type Export struct {
	data      []string
	timestamp time.Time
}

// Items returns a slice of all blocklist items.
func (e Export) Items() []string {
	return e.data
}

// Count returns the count of all items in the blocklist.
func (e Export) Count() int {
	return len(e.data)
}

// Timestamp returns the UTC timestamp when the blocklist
// was generated.
func (e Export) Timestamp() time.Time {
	return e.timestamp
}

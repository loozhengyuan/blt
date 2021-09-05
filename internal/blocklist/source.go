package blocklist

import (
	"fmt"
	"net/http"
)

// Source is the interface that describes an external source.
type Source interface {
	// Items returns a sorted list of all IPs or FQDNs from
	// the source.
	Items() ([]string, error)
}

// URLSource represents an external URL source.
type URLSource struct {
	url string
	p   Parser
}

var _ Source = (*URLSource)(nil)

// Items returns a sorted list of all IPs or FQDNs from
// the source.
func (r *URLSource) Items() ([]string, error) {
	resp, err := http.Get(r.url)
	if err != nil {
		return nil, fmt.Errorf("get request: %w", err)
	}
	defer resp.Body.Close()
	if s := resp.StatusCode; s != 200 {
		return nil, fmt.Errorf("non-200 http response: %v", s)
	}
	return r.p.Parse(resp.Body)
}

// NewURLSource returns the pointer to a new URLSource.
func NewURLSource(url string, p Parser) (*URLSource, error) {
	return &URLSource{
		url: url,
	}, nil
}

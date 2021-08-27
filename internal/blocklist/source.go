package blocklist

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
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
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}
	p, err := regexp.Compile(`(?m)^([^\s#]+)`)
	if err != nil {
		return nil, fmt.Errorf("compile regexp: %w", err)
	}
	return p.FindAllString(string(b), -1), nil
}

// NewURLSource returns the pointer to a new URLSource.
func NewURLSource(url string) (*URLSource, error) {
	return &URLSource{
		url: url,
	}, nil
}

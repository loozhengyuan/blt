package blocklist

import (
	"fmt"
	"io"
	"regexp"
)

const (
	REGEXP_IP   = `(?m)^([^\s#]+)`
	REGEXP_FQDN = `(?m)^([^\s#]+)`
)

// Parser is the interface that returns the list of parsed items.
type Parser interface {
	// Parse reads data from r and returns all parsed items
	// matching the pattern.
	//
	// A nil response indicates that no matches were found.
	Parse(r io.Reader) ([]string, error)

	// ParseString is like Parse but reads data from a string input.
	ParseString(s string) []string
}

// IPParser is a parser that returns matched patterns of
// IP addresses.
type IPParser struct {
	rg *regexp.Regexp
}

var _ Parser = (*IPParser)(nil)

func (p *IPParser) Parse(r io.Reader) ([]string, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read data: %w", err)
	}
	return p.ParseString(string(b)), nil
}

func (p *IPParser) ParseString(s string) []string {
	return p.rg.FindAllString(s, -1)
}

// NewIPParser returns the pointer to a new IPParser.
func NewIPParser() (*IPParser, error) {
	rg, err := regexp.Compile(REGEXP_IP)
	if err != nil {
		return nil, fmt.Errorf("compile regex: %w", err)
	}
	return &IPParser{rg: rg}, nil
}

// FQDNParser is a parser that returns matched patterns of
// fully-qualified domain names.
type FQDNParser struct {
	rg *regexp.Regexp
}

var _ Parser = (*FQDNParser)(nil)

func (p *FQDNParser) Parse(r io.Reader) ([]string, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read data: %w", err)
	}
	return p.ParseString(string(b)), nil
}

func (p *FQDNParser) ParseString(s string) []string {
	return p.rg.FindAllString(s, -1)
}

// NewFQDNParser returns the pointer to a new FQDNParser.
func NewFQDNParser() (*FQDNParser, error) {
	rg, err := regexp.Compile(REGEXP_FQDN)
	if err != nil {
		return nil, fmt.Errorf("compile regex: %w", err)
	}
	return &FQDNParser{rg: rg}, nil
}

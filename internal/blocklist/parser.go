package blocklist

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

const (
	REGEXP_IP   = `^([^\s]+)`
	REGEXP_FQDN = `([a-z0-9]+(-[a-z0-9]+)*\.)+[a-z]{2,}`
)

// Parser is the interface that returns the list of parsed items.
type Parser interface {
	// Parse reads data from r and returns all parsed items
	// matching the pattern.
	Parse(r io.Reader) ([]string, error)

	// ParseString is like Parse but reads data from a string input.
	ParseString(s string) ([]string, error)
}

// IPParser is a parser that returns matched patterns of
// IP addresses.
type IPParser struct {
	rg *regexp.Regexp
}

var _ Parser = (*IPParser)(nil)

func (p *IPParser) Parse(r io.Reader) ([]string, error) {
	m := make([]string, 0)
	s := bufio.NewScanner(r)
	for s.Scan() {
		// NOTE: The `regexp` library does not support positive
		// or negative lookaheads, hence the need to extract
		// non-commented strings before parsing with `regexp`.
		r := strings.Split(s.Text(), "#")[0]
		m = append(m, p.rg.FindAllString(r, -1)...)
	}
	return m, nil
}

func (p *IPParser) ParseString(s string) ([]string, error) {
	return p.Parse(strings.NewReader(s))
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
	m := make([]string, 0)
	s := bufio.NewScanner(r)
	for s.Scan() {
		// NOTE: The `regexp` library does not support positive
		// or negative lookaheads, hence the need to extract
		// non-commented strings before parsing with `regexp`.
		r := strings.Split(s.Text(), "#")[0]
		m = append(m, p.rg.FindAllString(r, -1)...)
	}
	return m, nil
}

func (p *FQDNParser) ParseString(s string) ([]string, error) {
	return p.Parse(strings.NewReader(s))
}

// NewFQDNParser returns the pointer to a new FQDNParser.
func NewFQDNParser() (*FQDNParser, error) {
	rg, err := regexp.Compile(REGEXP_FQDN)
	if err != nil {
		return nil, fmt.Errorf("compile regex: %w", err)
	}
	return &FQDNParser{rg: rg}, nil
}

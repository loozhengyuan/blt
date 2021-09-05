package blocklist

import (
	"reflect"
	"strings"
	"testing"
)

func TestIPParser_Parse(t *testing.T) {
	cases := map[string]struct {
		input string
		want  []string
	}{
		"empty": {
			input: "",
			want:  []string{},
		},
		"match_multiline_ipv4": {
			input: "192.168.1.1\n172.16.1.1\n10.1.1.1",
			want: []string{
				"192.168.1.1",
				"172.16.1.1",
				"10.1.1.1",
			},
		},
		"match_multiline_ipv6": {
			input: "fd12:3456:789a:1::1\nfd12:3456:789a:1::63",
			want: []string{
				"fd12:3456:789a:1::1",
				"fd12:3456:789a:1::63",
			},
		},
		"match_simple_single_ipv4": {
			input: "127.0.0.1",
			want: []string{
				"127.0.0.1",
			},
		},
		"match_simple_single_ipv6": {
			input: "::1",
			want: []string{
				"::1",
			},
		},
		"ignore_comments_line_ipv4": {
			input: "# 127.0.0.1",
			want:  []string{},
		},
		"ignore_comments_line_ipv6": {
			input: "# ::1",
			want:  []string{},
		},
		"ignore_comments_inline_ipv4": {
			input: "127.0.0.1 # 127.0.0.1",
			want: []string{
				"127.0.0.1",
			},
		},
		"ignore_comments_inline_ipv6": {
			input: "::1 # ::1",
			want: []string{
				"::1",
			},
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			p, err := NewIPParser()
			if err != nil {
				t.Fatalf("failed to init parser: %v", err)
			}
			got, err := p.Parse(strings.NewReader(tc.input))
			if err != nil {
				t.Fatalf("failed to parse input: %v", err)
			}
			if g, w := got, tc.want; !reflect.DeepEqual(g, w) {
				t.Errorf("output mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}

func TestFQDNParser_Parse(t *testing.T) {
	cases := map[string]struct {
		input string
		want  []string
	}{
		"empty": {
			input: "",
			want:  []string{},
		},
		"match_multiline": {
			input: "domain1.tld\ndomain2.tld",
			want: []string{
				"domain1.tld",
				"domain2.tld",
			},
		},
		"match_simple_single": {
			input: "domain1.tld",
			want: []string{
				"domain1.tld",
			},
		},
		"match_simple_multiple": {
			input: "domain1.tld domain2.tld",
			want: []string{
				"domain1.tld",
				"domain2.tld",
			},
		},
		"match_hosts_single": {
			input: "127.0.0.1 domain1.tld",
			want: []string{
				"domain1.tld",
			},
		},
		"match_hosts_multiple": {
			input: "127.0.0.1 domain1.tld domain2.tld",
			want: []string{
				"domain1.tld",
				"domain2.tld",
			},
		},
		"match_hosts_ipv4_prefix": {
			input: "127.0.0.1 domain1.tld",
			want: []string{
				"domain1.tld",
			},
		},
		"match_hosts_ipv6_prefix": {
			input: "::1 domain1.tld",
			want: []string{
				"domain1.tld",
			},
		},
		"ignore_comments_line": {
			input: "# domain1.tld",
			want:  []string{},
		},
		"ignore_comments_inline": {
			input: "domain1.tld # domain2.tld",
			want: []string{
				"domain1.tld",
			},
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			p, err := NewFQDNParser()
			if err != nil {
				t.Fatalf("failed to init parser: %v", err)
			}
			got, err := p.Parse(strings.NewReader(tc.input))
			if err != nil {
				t.Fatalf("failed to parse input: %v", err)
			}
			if g, w := got, tc.want; !reflect.DeepEqual(g, w) {
				t.Errorf("output mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}

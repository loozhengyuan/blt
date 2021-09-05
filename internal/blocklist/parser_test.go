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
			want:  nil, // nil response if no match
		},
		"ignore_comments_line_ipv4": {
			input: "# 127.0.0.1",
			want:  nil, // nil response if no match
		},
		"ignore_comments_line_ipv6": {
			input: "# ::1",
			want:  nil, // nil response if no match
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
			want:  nil, // nil response if no match
		},
		"ignore_comments_line": {
			input: "# domain1.tld",
			want:  nil, // nil response if no match
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

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
		"ignore_comments_ipv4": {
			input: "# 127.0.0.1",
			want:  nil, // nil response if no match
		},
		"ignore_comments_ipv6": {
			input: "# ::1",
			want:  nil, // nil response if no match
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
		"ignore_comments": {
			input: "# domain1.tld",
			want:  nil, // nil response if no match
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

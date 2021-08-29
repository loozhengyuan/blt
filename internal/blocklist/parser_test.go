package blocklist

import (
	"reflect"
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
			if g, w := p.ParseString(tc.input), tc.want; !reflect.DeepEqual(g, w) {
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
			input: "# google.com",
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
			if g, w := p.ParseString(tc.input), tc.want; !reflect.DeepEqual(g, w) {
				t.Errorf("output mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}

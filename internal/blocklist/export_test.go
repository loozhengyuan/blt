package blocklist

import (
	"reflect"
	"testing"
	"time"
)

func TestExport_Items(t *testing.T) {
	cases := map[string]struct {
		export Export
		want   []string
	}{
		"default": {
			export: Export{
				data: []string{"a", "b", "c"},
			},
			want: []string{"a", "b", "c"},
		},
		"empty": {
			export: Export{
				data: []string{},
			},
			want: []string{},
		},
		"nil": {
			want: nil,
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if g, w := tc.export.Items(), tc.want; !reflect.DeepEqual(g, w) {
				t.Errorf("output mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}

func TestExport_Count(t *testing.T) {
	cases := map[string]struct {
		export Export
		want   int
	}{
		"default": {
			export: Export{
				data: []string{"a", "b", "c"},
			},
			want: 3,
		},
		"empty": {
			export: Export{
				data: []string{},
			},
			want: 0,
		},
		"nil": {
			want: 0,
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if g, w := tc.export.Count(), tc.want; g != w {
				t.Errorf("output mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}

func TestExport_Timestamp(t *testing.T) {
	cases := map[string]struct {
		export Export
		want   time.Time
	}{
		"default": {
			export: Export{
				timestamp: time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		"nil": {
			want: time.Time{},
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if g, w := tc.export.Timestamp(), tc.want; g != w {
				t.Errorf("output mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}

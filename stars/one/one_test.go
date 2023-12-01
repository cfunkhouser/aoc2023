package one

import (
	"bytes"
	"testing"
)

func TestLineValue(t *testing.T) {
	type test struct {
		line line
		want int
	}
	for tn, tc := range map[string]test{
		"zero":            {},
		"single digit":    {"treb7uchet", 77},
		"multiple digits": {"a1b2c3d4e5f", 15},
	} {
		func(t *testing.T, tc *test) {
			t.Run(tn, func(t *testing.T) {
				t.Parallel()
				if got := tc.line.Value(); got != tc.want {
					t.Errorf("Value() mismatch: got: %d want: %d", got, tc.want)
				}
			})
		}(t, &tc)
	}
}

func BenchmarkLineValue(b *testing.B) {
	l := line("a1b2c3d4e5f")
	for i := 0; i < b.N; i++ {
		_ = l.Value()
	}
}

func TestFromDocument(t *testing.T) {
	type test struct {
		doc  string
		want int
	}
	for tn, tc := range map[string]test{
		"zero": {},
		"example from problem": {
			doc: `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`,
			want: 142,
		},
	} {
		func(t *testing.T, tc *test) {
			t.Run(tn, func(t *testing.T) {
				t.Parallel()
				buf := bytes.NewBufferString(tc.doc)
				if got := FromDocument(buf); got != tc.want {
					t.Errorf("FromDocument() mismatch: got: %d want: %d", got, tc.want)
				}
			})
		}(t, &tc)
	}
}

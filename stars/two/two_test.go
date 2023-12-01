package two

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
		"zero":                               {},
		"single digit":                       {"treb7uchet", 77},
		"multiple digits":                    {"a1b2c3d4e5f", 15},
		"mixed digits and spelled":           {"two1nine", 29},
		"only spelled":                       {"eightwothree", 83},
		"mixed digits, spelled, and garbage": {"abcone2threexyz", 13},
		"teens are tricky but we're smarter": {"7pqrstsixteen", 76},
		"overlaps are also tricky":           {"2zoneight", 28},
	} {
		func(t *testing.T, tc *test) {
			t.Run(tn, func(t *testing.T) {
				if got := tc.line.Value(); got != tc.want {
					t.Errorf("Value() mismatch: got: %d want: %d", got, tc.want)
				}
			})
		}(t, &tc)
	}
}

func BenchmarkLineValue(b *testing.B) {
	l := line("2zoneight")
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
			doc: `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`,
			want: 281,
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

package eight

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	type test struct {
		line    string
		want    *Card
		wantErr bool
	}

	for tn, tc := range map[string]test{
		"zero": {wantErr: true},
		"valid": {
			line: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			want: &Card{1, 4},
		},
	} {
		func(t *testing.T, tn string, tc *test) {
			t.Run(tn, func(t *testing.T) {
				got, _ := Parse(tc.line)
				if diff := cmp.Diff(got, tc.want); diff != "" {
					t.Errorf("Parse(): mismatch (-got,+want):\n%v", diff)
				}
			})
		}(t, tn, &tc)
	}
}

func TestFromDocument(t *testing.T) {
	type test struct {
		doc  string
		want Cards
	}

	for tn, tc := range map[string]test{
		"zero": {},
		"example from problem": {
			`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`,
			Cards{{1, 4}, {2, 2}, {3, 2}, {4, 1}, {5, 0}, {6, 0}},
		},
	} {
		func(t *testing.T, tn string, tc *test) {
			t.Run(tn, func(t *testing.T) {

				got := FromDocument(bytes.NewBufferString(tc.doc))
				if diff := cmp.Diff(got, tc.want); diff != "" {
					t.Errorf("Parse(): mismatch (-got,+want):\n%v", diff)
				}
			})
		}(t, tn, &tc)
	}
}

func cardsForTesting(tb testing.TB, doc string) Cards {
	tb.Helper()
	buf := bytes.NewBufferString(doc)
	return FromDocument(buf)
}

func TestCardsCount(t *testing.T) {
	type test struct {
		cards Cards
		want  int
	}

	for tn, tc := range map[string]test{
		"zero": {},
		"example from problem": {
			cards: cardsForTesting(t,
				`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`),
			want: 30,
		},
	} {
		func(t *testing.T, tn string, tc *test) {
			t.Run(tn, func(t *testing.T) {
				if got := tc.cards.Count(); got != tc.want {
					t.Errorf("Count(): mismatch: got: %d want: %d", got, tc.want)
				}
			})
		}(t, tn, &tc)
	}
}

package seven

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	type test struct {
		line    string
		want    Card
		wantErr bool
	}

	for tn, tc := range map[string]test{
		"zero": {wantErr: true},
		"valid": {
			line: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			want: Card{
				Winning: []int{41, 48, 83, 86, 17},
				Have:    []int{83, 86, 6, 31, 17, 9, 48, 53},
			},
		},
	} {
		func(t *testing.T, tn string, tc *test) {
			t.Run(tn, func(t *testing.T) {
				got := Parse(tc.line)
				if diff := cmp.Diff(got, tc.want); diff != "" {
					t.Errorf("Parse(): mismatch (-got,+want):\n%v", diff)
				}
			})
		}(t, tn, &tc)
	}
}

func TestCardMatches(t *testing.T) {
	type test struct {
		card Card
		want []int
	}

	for tn, tc := range map[string]test{
		"zero": {},
		"valid static": {
			card: Card{
				Winning: []int{41, 48, 83, 86, 17},
				Have:    []int{83, 86, 6, 31, 17, 9, 48, 53},
			},
			want: []int{17, 48, 83, 86},
		},
		"example card 1": {
			card: Parse("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"),
			want: []int{17, 48, 83, 86},
		},
		"example card 2": {
			card: Parse("Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"),
			want: []int{32, 61},
		},
		"example card 3": {
			card: Parse("Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1"),
			want: []int{1, 21},
		},
		"example card 4": {
			card: Parse("Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83"),
			want: []int{84},
		},
		"example card 5": {
			card: Parse("Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36"),
		},
		"example card 6": {
			card: Parse("Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"),
		},
	} {
		func(t *testing.T, tn string, tc *test) {
			t.Run(tn, func(t *testing.T) {
				got := tc.card.Matches()
				if diff := cmp.Diff(got, tc.want); diff != "" {
					t.Errorf("Matches(): mismatch (-got,+want):\n%v", diff)
				}
			})
		}(t, tn, &tc)
	}
}

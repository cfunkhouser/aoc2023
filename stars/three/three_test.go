package three

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseGame(t *testing.T) {
	type test struct {
		line string
		want *Game
	}

	for tn, tc := range map[string]test{
		"empty line produces zero value": {"", &Game{}},
		"valid game produces valid Game": {
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			&Game{4, 14, 3, 15},
		},
	} {
		func(t *testing.T, tn string, tc *test) {
			t.Run(tn, func(t *testing.T) {
				got := ParseGame(tc.line)
				if diff := cmp.Diff(got, tc.want); diff != "" {
					t.Errorf("ParseGame(): mismatch: (-got,+want):\n%v", diff)
				}
			})
		}(t, tn, &tc)
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
			`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`,
			8,
		},
	} {
		func(t *testing.T, tn string, tc *test) {
			t.Run(tn, func(t *testing.T) {
				buf := bytes.NewBufferString(tc.doc)
				if got := FromDocument(buf, 12, 13, 14); got != tc.want {
					t.Errorf("FromDocument(): mismatch: got %d want %d", got, tc.want)
				}
			})
		}(t, tn, &tc)
	}
}

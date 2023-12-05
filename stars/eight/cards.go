package eight

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/cfunkhouser/aoc2023/util"
)

type Card struct {
	ID      int
	Matches int
}

func (c *Card) String() string {
	if c.Matches < 1 {
		return fmt.Sprintf("[%-5d]", c.ID)
	}
	return fmt.Sprintf("[%-2d+%2d]", c.ID, c.Matches)
}

var numsRe = regexp.MustCompile(`\d+`)

func Parse(s string) (*Card, error) {
	widx := strings.Index(s, ":")
	hidx := strings.Index(s, "|")
	if widx == -1 || hidx == -1 {
		return nil, errors.New("invalid Card")
	}

	ididxs := numsRe.FindStringIndex(s)
	id, err := strconv.Atoi(s[ididxs[0]:ididxs[1]])
	if err != nil {
		return nil, err
	}
	w := s[widx+1 : hidx]
	h := s[hidx+1:]

	var winning []int
	for _, wm := range numsRe.FindAllStringIndex(w, -1) {
		n, err := strconv.Atoi(w[wm[0]:wm[1]])
		if err != nil {
			return nil, fmt.Errorf("invalid Card: %w", err)
		}
		winning = append(winning, n)
	}

	var have []int
	for _, hm := range numsRe.FindAllStringIndex(h, -1) {
		n, err := strconv.Atoi(h[hm[0]:hm[1]])
		if err != nil {
			return nil, fmt.Errorf("invalid Card: %w", err)
		}
		have = append(have, n)
	}

	return &Card{
		ID:      id,
		Matches: len(util.NewSet(winning).Intersection(util.NewSet(have)).Values()),
	}, nil
}

// Cards is a collection of cards.
type Cards []*Card

func (cards Cards) String() string {
	var buf bytes.Buffer
	for _, c := range cards {
		if _, err := buf.WriteString(c.String()); err != nil {
			panic(err)
		}
	}
	return buf.String()
}

// Count expands the cards according the match rules and returns the total
// number.
func (cards Cards) Count() int {
	expanded := slices.Clone(cards)

	// Each card is counted once before copies are included.
	ncards := len(cards)
	count := make([]int, ncards)
	for i := 0; i < ncards; i++ {
		count[i] = 1
	}

	for i := 0; i < len(expanded); i++ {
		c := expanded[i]
		m := c.Matches
		copies := cards[c.ID : c.ID+m]
		for _, cc := range copies {
			count[cc.ID-1]++
		}
		expanded = append(expanded, copies...)
	}
	return util.Sum(count)
}

// FromDocument calculates the sum of point values for all scratch off cards.
func FromDocument(doc io.Reader) (cards Cards) {
	s := bufio.NewScanner(doc)
	for s.Scan() {
		card, err := Parse(s.Text())
		if err != nil {
			panic(err)
		}
		cards = append(cards, card)
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return
}

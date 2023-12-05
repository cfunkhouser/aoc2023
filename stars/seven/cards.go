package seven

import (
	"bufio"
	"io"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/cfunkhouser/aoc2023/util"
)

type Card struct {
	Winning []int
	Have    []int
}

var numsRe = regexp.MustCompile(`\d+`)

func Parse(s string) (ret Card) {
	widx := strings.Index(s, ":")
	hidx := strings.Index(s, "|")
	if widx == -1 || hidx == -1 {
		return ret
	}

	w := s[widx+1 : hidx]
	h := s[hidx+1:]
	for _, wm := range numsRe.FindAllStringIndex(w, -1) {
		n, err := strconv.Atoi(w[wm[0]:wm[1]])
		if err != nil {
			panic(err)
		}
		ret.Winning = append(ret.Winning, n)
	}
	for _, hm := range numsRe.FindAllStringIndex(h, -1) {
		n, err := strconv.Atoi(h[hm[0]:hm[1]])
		if err != nil {
			panic(err)
		}
		ret.Have = append(ret.Have, n)
	}
	return ret
}

func (c *Card) Matches() []int {
	return util.NewSet(c.Winning).Intersection(util.NewSet(c.Have)).Values()
}

// FromDocument calculates the sum of point values for all scratch off cards.
func FromDocument(doc io.Reader) (points int) {
	s := bufio.NewScanner(doc)
	for s.Scan() {
		card := Parse(s.Text())
		matches := len(card.Matches())
		if matches > 0 {
			points += int(math.Pow(2, float64(matches-1)))
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return
}

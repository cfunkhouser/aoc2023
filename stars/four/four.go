// Package four solves for the fourth star in Advent of Code 2023.
// See: https://adventofcode.com/2023/day/2
package four

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// Game represents a single round of the (truly terrible) Snow Island Elf game.
// The Min* values are the highest observed value of each color.
type Game struct {
	ID       int
	MinRed   int
	MinGreen int
	MinBlue  int
}

// Power of the set of cubes in a Game.
func (gg *Game) Power() int {
	if gg == nil {
		return 0
	}
	return gg.MinRed * gg.MinGreen * gg.MinBlue
}

func gameIDAndColorString(s string) (id int, colors string) {
	segs := strings.Split(s, ":")
	if len(segs) != 2 {
		return
	}
	var err error
	id, err = strconv.Atoi(strings.TrimPrefix(segs[0], "Game "))
	if err != nil {
		return
	}
	colors = segs[1]
	return
}

var colorsRe = regexp.MustCompile(`\s*(\d+)\s+(red|green|blue)\s*`)

func parseColors(s string) (r int, g int, b int) {
	for _, m := range colorsRe.FindAllStringSubmatch(s, -1) {
		if len(m) != 3 {
			continue
		}
		v, err := strconv.Atoi(m[1])
		if err != nil {
			panic(err)
		}
		switch color := m[2]; color {
		case "red":
			r = v
		case "green":
			g = v
		case "blue":
			b = v
		default:
			panic("this shouldn't be possible, but here we are")
		}
	}
	return
}

func maxColors(s string) (r int, g int, b int) {
	for _, seg := range strings.Split(s, ";") {
		lr, lg, lb := parseColors(seg)
		if lr > r {
			r = lr
		}
		if lg > g {
			g = lg
		}
		if lb > b {
			b = lb
		}
	}
	return
}

func ParseGame(s string) *Game {
	var ret Game
	var colors string
	ret.ID, colors = gameIDAndColorString(s)
	ret.MinRed, ret.MinGreen, ret.MinBlue = maxColors(colors)
	return &ret
}

// FromDocument calculates the sum of the IDs of all games which would have been
// possible given the values of r, g, and b.
func FromDocument(doc io.Reader) (value int) {
	s := bufio.NewScanner(doc)
	for s.Scan() {
		ss := s.Text()
		gg := ParseGame(ss)
		value += gg.Power()
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return
}

var (
	filePath string

	starCmd = &cobra.Command{
		Use:     "four",
		Aliases: []string{"fourth", "4"},
		Short:   "Calculate the sum of the power of each minimal set.",
		Long: `Calculate the sum of the power of each minimal set.
		
	If no value is provided for -f / --file the document is read from STDIN.
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			f := os.Stdin
			if filePath != "" {
				var err error
				if f, err = os.Open(filePath); err != nil {
					return err
				}
				defer f.Close()
			}
			fmt.Println(FromDocument(f))
			return nil
		},
	}
)

func init() {
	starCmd.Flags().StringVarP(&filePath, "file", "f", "",
		"Path to the trebuchet calibration document. Optional.")
}

// RegisterOn the provided command.
func RegisterOn(cmd *cobra.Command) {
	cmd.AddCommand(starCmd)
}

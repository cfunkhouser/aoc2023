// Package three solves for the third star in Advent of Code 2023.
// See: https://adventofcode.com/2023/day/2
package three

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
// The Max* values are the highest observed value of each color.
type Game struct {
	ID       int
	MaxRed   int
	MaxGreen int
	MaxBlue  int
}

// Possible is true if the number of cubes of each color in the bag for a game
// could possibly be the number specified.
func (gg *Game) Possible(r, g, b int) bool {
	return r >= gg.MaxRed && g >= gg.MaxGreen && b >= gg.MaxBlue
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
	ret.MaxRed, ret.MaxGreen, ret.MaxBlue = maxColors(colors)
	return &ret
}

// FromDocument calculates the sum of the IDs of all games which would have been
// possible given the values of r, g, and b.
func FromDocument(doc io.Reader, r, g, b int) (value int) {
	s := bufio.NewScanner(doc)
	for s.Scan() {
		ss := s.Text()
		gg := ParseGame(ss)
		if gg.Possible(r, g, b) {
			value += gg.ID
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return
}

var (
	filePath         string
	red, green, blue int

	starCmd = &cobra.Command{
		Use:     "three",
		Aliases: []string{"third", "3"},
		Short:   "Calculate the sum of the IDs of possible games for given RGB values.",
		Long: `Calculate the sum of the IDs of possible games for given RGB values.
		
	If no value is provided for -f / --file the document is read from STDIN.
	Override the RGB values with the -r, -g, and -b flags, respectively.
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
			fmt.Println(FromDocument(f, red, green, blue))
			return nil
		},
	}
)

func init() {
	starCmd.Flags().StringVarP(&filePath, "file", "f", "",
		"Path to the trebuchet calibration document. Optional.")

	starCmd.Flags().IntVarP(&red, "red", "r", 12, "Red value to check.")
	starCmd.Flags().IntVarP(&green, "green", "g", 13, "Green value to check.")
	starCmd.Flags().IntVarP(&blue, "blue", "b", 14, "Blue value to check.")
}

// RegisterOn the provided command.
func RegisterOn(cmd *cobra.Command) {
	cmd.AddCommand(starCmd)
}

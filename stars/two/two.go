// Package two solves for the second star in Advent of Code 2023.
// See: https://adventofcode.com/2023/day/1
package two

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
	"strconv"

	"github.com/spf13/cobra"
)

// line represents a single line of input in the trebuchet calibration document.
type line string

var (
	matchers multiRegexp = []*regexp.Regexp{
		regexp.MustCompile(`\d`), // digits

		// The regexp package uses re2 syntax, which does not support lookahead.
		// To work around this, multiple regular expressions are used, and the
		// results are merged.
		regexp.MustCompile(`one`),
		regexp.MustCompile(`two`),
		regexp.MustCompile(`three`),
		regexp.MustCompile(`four`),
		regexp.MustCompile(`five`),
		regexp.MustCompile(`six`),
		regexp.MustCompile(`seven`),
		regexp.MustCompile(`eight`),
		regexp.MustCompile(`nine`),
	}

	words = map[string]byte{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}
)

type multiRegexp []*regexp.Regexp

// FindAllStringIndex aggregates the results of calls to multiple regular
// expressions' FindAllStringIndex. The return value is sorted.
func (r *multiRegexp) FindAllStringIndex(s string, n int) (result [][]int) {
	for _, re := range *r {
		result = append(result, re.FindAllStringIndex(s, n)...)
	}
	slices.SortStableFunc(result, func(l, r []int) int {
		if l[0] == r[0] {
			// If the starting index matches, compare the terminal index.
			return l[1] - r[1]
		}
		return l[0] - r[0]
	})
	return
}

// Value extracts the numerical value from the given calibration value.
func (v line) Value() int {
	idxs := matchers.FindAllStringIndex(string(v), -1)
	n := len(idxs)
	if n == 0 {
		return 0
	}

	b := make([]byte, 2)
	vs := string(v)

	if (idxs[0][1] - idxs[0][0]) > 1 {
		b[0] = words[vs[idxs[0][0]:idxs[0][1]]]
	} else {
		b[0] = vs[idxs[0][0]]
	}

	if (idxs[n-1][1] - idxs[n-1][0]) > 1 {
		b[1] = words[vs[idxs[n-1][0]:idxs[n-1][1]]]
	} else {
		b[1] = vs[idxs[n-1][0]]
	}

	vi, err := strconv.Atoi(string(b))
	if err != nil {
		panic(err)
	}
	return vi
}

// FromDocument calculates the overall calibration value from a calibration
// document containing one value per line.
func FromDocument(r io.Reader) (value int) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		value += line(s.Text()).Value()
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return
}

var filePath string

func init() {
	starCmd.Flags().StringVarP(&filePath, "file", "f", "",
		"Path to the trebuchet calibration document. Optional.")
}

var starCmd = &cobra.Command{
	Use:     "two",
	Aliases: []string{"second", "2"},
	Short:   "Calculate value from a trebuchet calibration document",
	Long: `Calculate the overall trebuchet calibration value from a calibration document.
	
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

// RegisterOn the provided command.
func RegisterOn(cmd *cobra.Command) {
	cmd.AddCommand(starCmd)
}

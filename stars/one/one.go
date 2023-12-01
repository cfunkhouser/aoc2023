// Package one solves for the first star in Advent of Code 2023.
// See: https://adventofcode.com/2023/day/1
package one

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
)

// line represents a single line of input in the trebuchet calibration document.
type line string

var digitRe = regexp.MustCompile(`\d`)

// Value extracts the numerical value from the given calibration value.
func (v line) Value() int {
	idxs := digitRe.FindAllStringIndex(string(v), -1)
	n := len(idxs)
	if n == 0 {
		return 0
	}
	vi, err := strconv.Atoi(string([]byte{v[idxs[0][0]], v[idxs[n-1][0]]}))
	if err != nil {
		panic(err)
	}
	return vi
}

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
	Use:     "one",
	Aliases: []string{"first", "1"},
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

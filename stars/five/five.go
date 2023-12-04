// Package five solves for the fifth star in Advent of Code 2023.
// See: https://adventofcode.com/2023/day/3
package five

import (
	"fmt"
	"os"

	"github.com/cfunkhouser/aoc2023/gondola"
	"github.com/cfunkhouser/aoc2023/util"
	"github.com/spf13/cobra"
)

var (
	filePath string

	starCmd = &cobra.Command{
		Use:     "five",
		Aliases: []string{"fifth", "5"},
		Short:   "Calculate the sum of part numbers in a gondola schematic.",
		Long: `Calculate the sum of part numbers in a gondola schematic.
		
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
			fmt.Println(util.Sum(gondola.FromDocument(f).PartNumbers()))
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

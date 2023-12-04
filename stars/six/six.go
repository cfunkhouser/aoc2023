// Package six solves for the sixth star in Advent of Code 2023.
// See: https://adventofcode.com/2023/day/3
package six

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
		Use:     "six",
		Aliases: []string{"sixth", "6"},
		Short:   "Calculate the sum of gear ratios from a gondola schematic.",
		Long: `Calculate the sum of gear ratios from a gondola schematic.
		
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
			fmt.Println(util.Sum(gondola.FromDocument(f).GearRatios()))
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

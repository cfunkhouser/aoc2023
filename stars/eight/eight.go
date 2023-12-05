// Package eight solves for the eighth star in Advent of Code 2023.
// See: https://adventofcode.com/2023/day/4
package eight

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	filePath string

	starCmd = &cobra.Command{
		Use:     "eight",
		Aliases: []string{"eighth", "8"},
		Short:   "Calculate the total number of scratch off cards.",
		Long: `Calculate the total number of scratch off cards.

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
			fmt.Println(FromDocument(f).Count())
			return nil
		},
	}
)

func init() {
	starCmd.Flags().StringVarP(&filePath, "file", "f", "",
		"Path to the scratch card values. Optional.")
}

// RegisterOn the provided command.
func RegisterOn(cmd *cobra.Command) {
	cmd.AddCommand(starCmd)
}

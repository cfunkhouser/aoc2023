// Binary aoc2023 implements solutions for Advent of Code 2023.
package main

import (
	"os"

	"github.com/cfunkhouser/aoc2023/stars"
	"github.com/spf13/cobra"
)

var (
	// rootCmd represents the base command when called without any subcommands.
	rootCmd = &cobra.Command{
		Use:   "aoc2023",
		Short: "Advent of Code 2023",
		Long:  "Advent of Code 2023",
	}
)

func main() {
	stars.RegisterOn(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// Package stars exposes the solution to each star collection solution for
// aoc2023.
package stars

import (
	"github.com/spf13/cobra"

	"github.com/cfunkhouser/aoc2023/stars/one"
)

var starCmd = &cobra.Command{
	Use:   "star",
	Short: "Solve for an AoC 2023 Star",
	Long:  "Solve for an AoC 2023 Star",
}

func init() {
	one.RegisterOn(starCmd)
}

// RegisterOn the provided command.
func RegisterOn(cmd *cobra.Command) {
	cmd.AddCommand(starCmd)
}

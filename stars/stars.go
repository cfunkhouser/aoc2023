// Package stars exposes the solution to each star collection solution for
// aoc2023.
package stars

import (
	"github.com/spf13/cobra"

	"github.com/cfunkhouser/aoc2023/stars/five"
	"github.com/cfunkhouser/aoc2023/stars/four"
	"github.com/cfunkhouser/aoc2023/stars/one"
	"github.com/cfunkhouser/aoc2023/stars/seven"
	"github.com/cfunkhouser/aoc2023/stars/six"
	"github.com/cfunkhouser/aoc2023/stars/three"
	"github.com/cfunkhouser/aoc2023/stars/two"
)

var starCmd = &cobra.Command{
	Use:   "star",
	Short: "Solve for an AoC 2023 Star",
	Long:  "Solve for an AoC 2023 Star",
}

func init() {
	one.RegisterOn(starCmd)
	two.RegisterOn(starCmd)
	three.RegisterOn(starCmd)
	four.RegisterOn(starCmd)
	five.RegisterOn(starCmd)
	six.RegisterOn(starCmd)
	seven.RegisterOn(starCmd)
}

// RegisterOn the provided command.
func RegisterOn(cmd *cobra.Command) {
	cmd.AddCommand(starCmd)
}

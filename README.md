# Advent of Code 2023

These are my solutions for the [Advent of Code
2023](https://adventofcode.com/2023), all wrapped up in a single command line
tool written in Go.

I will update it whenever I have an answer for a given challenge. If a challenge
is missing, you may assume I didn't solve it. This tool may be used for anything
you may desire.

## Install

The tool can be installed by running:

`go install github.com/cfunkhouser/aoc2023/cmd/aoc2023@latest`

## Running

Each completed challenge earns the challenger a star. To solve the problem and
earn a star, run `aoc2023 star $STAR`, where `$STAR` is the number of the
challenge. For more details about what each challenge expects as input, you may
run `aoc2023 star help $STAR`.

The full help output is:

```
$ aoc2023 help star
Solve for an AoC 2023 Star

Usage:
  aoc2023 star [command]

Available Commands:
  four        Calculate the sum of the power of each minimal set.
  one         Calculate value from a trebuchet calibration document
  three       Calculate the sum of the IDs of possible games for given RGB values.
  two         Calculate value from a trebuchet calibration document

Flags:
  -h, --help   help for star

Use "aoc2023 star [command] --help" for more information about a command.
```
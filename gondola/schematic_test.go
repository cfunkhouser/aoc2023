package gondola

import (
	"bytes"
	"testing"

	"github.com/cfunkhouser/aoc2023/util"
	"github.com/google/go-cmp/cmp"

	_ "embed"
)

func TestPointAdjacent(t *testing.T) {
	type test struct {
		p    point
		want Adjacent[point]
	}

	for tn, tc := range map[string]test{
		"zero": {
			want: Adjacent[point]{
				{-1, -1},
				{0, -1},
				{1, -1},
				{1, 0},
				{1, 1},
				{0, 1},
				{-1, 1},
				{-1, 0},
			},
		},
	} {
		func(t *testing.T, tn string, tc *test) {
			t.Run(tn, func(t *testing.T) {
				got := tc.p.Adjacent()
				if diff := cmp.Diff(got, tc.want); diff != "" {
					t.Errorf("Adjacent(): mismatch (-got,+want):\n%v", diff)
				}
			})
		}(t, tn, &tc)
	}
}

func TestValueNumber(t *testing.T) {
	type test struct {
		val    Value
		want   int
		wantOk bool
	}

	for tn, tc := range map[string]test{
		"zero":            {},
		"nonzero runic":   {Value{r: util.Pointy[rune]('r')}, 0, false},
		"nonzero numeric": {Value{n: util.Pointy[int](42)}, 42, true},
	} {
		func(t *testing.T, tn string, tc *test) {
			t.Run(tn, func(t *testing.T) {
				got, ok := tc.val.Number()
				if ok != tc.wantOk {
					t.Errorf("Number(): ok mismatch: got: %v want: %v", ok, tc.wantOk)
				}
				if got != tc.want {
					t.Errorf("Number(): mismatch: got: %v want: %v", got, tc.want)
				}
			})
		}(t, tn, &tc)
	}
}

func TestValueRune(t *testing.T) {
	type test struct {
		val    Value
		want   rune
		wantOk bool
	}

	for tn, tc := range map[string]test{
		"zero":            {},
		"nonzero runic":   {Value{r: util.Pointy[rune]('r')}, 'r', true},
		"nonzero numeric": {Value{n: util.Pointy[int](42)}, 0, false},
	} {
		func(t *testing.T, tn string, tc *test) {
			t.Run(tn, func(t *testing.T) {
				got, ok := tc.val.Rune()
				if ok != tc.wantOk {
					t.Errorf("Rune(): ok mismatch: got: %v want: %v", ok, tc.wantOk)
				}
				if got != tc.want {
					t.Errorf("Rune(): mismatch: got: %v want: %v", got, tc.want)
				}
			})
		}(t, tn, &tc)
	}
}

func TestValueString(t *testing.T) {
	type test struct {
		val  *Value
		want string
	}

	for tn, tc := range map[string]test{
		"nil":             {nil, "."},
		"zero":            {&Value{}, "."},
		"nonzero runic":   {&Value{r: util.Pointy[rune]('r')}, "r"},
		"nonzero numeric": {&Value{n: util.Pointy[int](42)}, "42"},
	} {
		func(t *testing.T, tn string, tc *test) {
			t.Run(tn, func(t *testing.T) {
				if got := tc.val.String(); got != tc.want {
					t.Errorf("String(): mismatch: got: %v want: %v", got, tc.want)
				}
			})
		}(t, tn, &tc)
	}
}

func TestRawSchematicCompact(t *testing.T) {
	type test struct {
		rs   rawSchematic
		want *Schematic
	}

	cmpOpts := []cmp.Option{
		cmp.AllowUnexported(Schematic{}),
		cmp.AllowUnexported(Value{}),
	}
	c42 := &Cell{Value: Value{n: util.Pointy[int](42)}}
	c7 := &Cell{Value: Value{n: util.Pointy[int](7)}}
	cx := &Cell{Value: Value{r: util.Pointy[rune]('x')}}

	for tn, tc := range map[string]test{
		"zero": {want: &Schematic{}},
		"all nil cells": {
			rs:   rawSchematic{{nil, nil}, {nil, nil}},
			want: &Schematic{},
		},
		"no links": {
			rs: rawSchematic{
				{c42, c42, nil},
				{nil, cx, nil},
				{nil, nil, c7},
			},
			want: &Schematic{
				Characters: []*Cell{cx},
			},
		},
	} {
		func(t *testing.T, tn string, tc *test) {
			t.Run(tn, func(t *testing.T) {
				got := tc.rs.Compact()
				if diff := cmp.Diff(got, tc.want, cmpOpts...); diff != "" {
					t.Errorf("Compact(): mismatch: (-got,+want):\n%v", diff)
				}
			})
		}(t, tn, &tc)
	}
}

func TestRawSchematicString(t *testing.T) {
	type test struct {
		rs   rawSchematic
		want string
	}

	c42 := &Cell{Value: Value{n: util.Pointy[int](42)}}
	c7 := &Cell{Value: Value{n: util.Pointy[int](7)}}
	cx := &Cell{Value: Value{r: util.Pointy[rune]('x')}}

	for tn, tc := range map[string]test{
		"zero": {},
		"all nil cells": {
			rs:   rawSchematic{{nil, nil}, {nil, nil}},
			want: "..\n..",
		},
		"no links": {
			rs: rawSchematic{
				{c42, c42, nil},
				{nil, cx, nil},
				{nil, nil, c7},
			},
			want: "42.\n.x.\n..7",
		},
	} {
		func(t *testing.T, tn string, tc *test) {
			t.Run(tn, func(t *testing.T) {
				got := tc.rs.String()
				if diff := cmp.Diff(got, tc.want); diff != "" {
					t.Errorf("Compact(): mismatch: (-got,+want):\n%v", diff)
				}
			})
		}(t, tn, &tc)
	}
}

func TestRawSchematicStringIdentity(t *testing.T) {
	type test struct {
		doc string
	}

	for tn, tc := range map[string]test{
		"empty": {},
		"example from problem": {`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`},
	} {
		func(t *testing.T, tn string, tc *test) {
			t.Run(tn, func(t *testing.T) {
				got := rawFromDocument(bytes.NewBufferString(tc.doc)).String()
				if diff := cmp.Diff(got, tc.doc); diff != "" {
					t.Errorf("rawDocument Identity: mismatch (-got,+want):\n%v", diff)
				}
			})
		}(t, tn, &tc)
	}
}

func TestSchematicPartNumbers(t *testing.T) {
	type test struct {
		doc  string
		want []int
	}

	for tn, tc := range map[string]test{
		"empty": {},
		"example from problem": {
			`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
			[]int{35, 467, 592, 598, 617, 633, 664, 755}, // sorted
		},
		"example from problem with one fewer columns": {
			`467..114.
...*.....
..35..633
......#..
617*.....
.....+.58
..592....
......755
...$.*...
.664.598.`,
			[]int{35, 467, 592, 598, 617, 633, 664, 755}, // sorted
		},
	} {
		func(t *testing.T, tn string, tc *test) {
			t.Run(tn, func(t *testing.T) {
				got := FromDocument(bytes.NewBufferString(tc.doc)).PartNumbers()
				if diff := cmp.Diff(got, tc.want); diff != "" {
					t.Errorf("PartNumbers(): mismatch (-got,+want):\n%v", diff)
				}
			})
		}(t, tn, &tc)
	}
}

func TestSchematicGearRatios(t *testing.T) {
	type test struct {
		doc  string
		want []int
	}

	for tn, tc := range map[string]test{
		"empty": {},
		"example from problem": {
			`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
			[]int{16345, 451490}, // sorted
		},
	} {
		func(t *testing.T, tn string, tc *test) {
			t.Run(tn, func(t *testing.T) {
				got := FromDocument(bytes.NewBufferString(tc.doc)).GearRatios()
				if diff := cmp.Diff(got, tc.want); diff != "" {
					t.Errorf("GearRatios(): mismatch (-got,+want):\n%v", diff)
				}
			})
		}(t, tn, &tc)
	}
}

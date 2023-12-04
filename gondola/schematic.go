package gondola

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"slices"
	"strconv"
	"text/tabwriter"

	"github.com/cfunkhouser/aoc2023/util"
)

// Value at a cell in a Schematic
type Value struct {
	n *int
	r *rune
}

// Number unwraps the numeric value in a Cell, if any.
func (v *Value) Number() (int, bool) {
	if v.n != nil {
		return *(v.n), true
	}
	return 0, false
}

// Rune unwraps the runic value in a Cell, if any.
func (v *Value) Rune() (rune, bool) {
	if v.r != nil {
		return *(v.r), true
	}
	return 0, false
}

// String produces a string representation of a Value.
func (v *Value) String() string {
	if v == nil {
		return "."
	}
	if n, ok := v.Number(); ok {
		return strconv.Itoa(n)
	}
	if r, ok := v.Rune(); ok {
		return string([]rune{r})
	}
	return "."
}

// Cell in a Schematic.
type Cell struct {
	Value
	Adjacent Adjacent[*Cell]
}

// String produces a string representation of a Cell.
func (c *Cell) String() string {
	if c == nil {
		return "."
	}
	return c.Value.String()
}

// DebugString renders the cell and adjacent cells graphically for debugging.
func (c *Cell) DebugString() string {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 2, 1, 1, ' ', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\n", c.Adjacent[NW], c.Adjacent[N], c.Adjacent[NE])
	fmt.Fprintf(w, "%s\t%s\t%s\n", c.Adjacent[W], c, c.Adjacent[E])
	fmt.Fprintf(w, "%s\t%s\t%s\n", c.Adjacent[SW], c.Adjacent[S], c.Adjacent[SE])
	w.Flush()
	return buf.String()
}

// ValidAdjacent returns a slice of all adjacent cells which are not nil.
func (c *Cell) ValidAjacent() (ret []*Cell) {
	for _, ac := range c.Adjacent {
		if ac != nil {
			ret = append(ret, ac)
		}
	}
	return
}

// Schematic of a gondola.
type Schematic struct {
	// Characters are the cells in the schematic containing characters.
	Characters []*Cell
}

// PartNumbers from the schematic. The resulting list is sorted.
func (s *Schematic) PartNumbers() (ret []int) {
	ncells := make(map[*Cell]int)
	// Hash the cells by pointer to prevent duplicate part numbers.
	for _, c := range s.Characters {
		for _, ac := range c.ValidAjacent() {
			if n, ok := ac.Number(); ok {
				ncells[ac] = n
			}
		}
	}
	for _, n := range ncells {
		ret = append(ret, n)
	}
	slices.SortStableFunc(ret, func(l, r int) int {
		return l - r
	})
	return
}

type point struct {
	X, Y int
}

// Direction from the reference point to seek an adjacent value.
type Direction int

const (
	NW Direction = iota
	N
	NE
	E
	SE
	S
	SW
	W
)

// Reverse relationship of the direction.
func (d Direction) Reverse() Direction {
	switch d {
	case NW:
		return SE
	case N:
		return S
	case NE:
		return SW
	case E:
		return W
	case SE:
		return NW
	case S:
		return N
	case SW:
		return NE
	case W:
		return E
	default:
		panic("invalid direction")
	}
}

// Adjacent holds values adjacent to a reference point on a two-dimensional
// cartesian plane. They are organized clockwise around the (unincluded)
// reference, starting with the Northwest value at index 0. Use the Direction
// constants for easy reference.
type Adjacent[T any] [8]T

func (p *point) Adjacent() Adjacent[point] {
	return Adjacent[point]{
		{p.X - 1, p.Y - 1}, // NW
		{p.X, p.Y - 1},     // N
		{p.X + 1, p.Y - 1}, // NE
		{p.X + 1, p.Y},     // E
		{p.X + 1, p.Y + 1}, // SE
		{p.X, p.Y + 1},     // S
		{p.X - 1, p.Y + 1}, // SW
		{p.X - 1, p.Y},     // W
	}
}

func (p *point) ValidWithin(extents point) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < extents.X && p.Y < extents.Y
}

type rawSchematic [][]*Cell

func (rs rawSchematic) linkAdjacent(loc point, cell *Cell) {
	var h, w int
	if h = len(rs); h == 0 {
		return
	}
	if w = len(rs[0]); w == 0 {
		return
	}

	for i, adj := range loc.Adjacent() {
		d := Direction(i)
		if adj.ValidWithin(point{w, h}) {
			ac := rs[adj.Y][adj.X]
			if ac == nil {
				continue
			}
			cell.Adjacent[d] = ac
			ac.Adjacent[d.Reverse()] = cell
		}
	}
}

// Compact a raw schematic into a usable Schematic.
func (rs rawSchematic) Compact() *Schematic {
	var ret Schematic
	for y, row := range rs {
		for x, cell := range row {
			if cell == nil {
				continue
			}
			if _, ok := cell.Rune(); ok {
				ret.Characters = append(ret.Characters, cell)
			}
			rs.linkAdjacent(point{x, y}, cell)
		}
	}
	return &ret
}

// String produces a string representation of a rawSchematic.
func (rs rawSchematic) String() string {
	var buf bytes.Buffer
	lri := len(rs) - 1
	for i, row := range rs {
		var last *Cell
		for _, cell := range row {
			if last != nil && last == cell {
				continue
			}
			last = cell
			fmt.Fprint(&buf, cell.String())
		}
		if i < lri {
			fmt.Fprintln(&buf)
		}
	}
	return buf.String()
}

var (
	numRe  = regexp.MustCompile(`\d+`)
	charRe = regexp.MustCompile(`[^\d.]`)
)

func lineToCells(l string) []*Cell {
	numMatches := numRe.FindAllStringIndex(l, -1)
	charMatches := charRe.FindAllStringIndex(l, -1)
	ret := make([]*Cell, len(l))
	for _, nm := range numMatches {
		n, err := strconv.Atoi(l[nm[0]:nm[1]])
		if err != nil {
			panic(err)
		}
		c := &Cell{
			Value: Value{
				n: util.Pointy(n),
			},
		}
		for i := nm[0]; i < nm[1]; i++ {
			ret[i] = c
		}
	}
	for _, cm := range charMatches {
		ret[cm[0]] = &Cell{
			Value: Value{
				r: util.Pointy(rune(l[cm[0]])),
			},
		}
	}
	return ret
}

func rawFromDocument(doc io.Reader) (ret rawSchematic) {
	s := bufio.NewScanner(doc)
	for s.Scan() {
		ret = append(ret, lineToCells(s.Text()))
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return
}

// FromDocument produces a Schematic from the contents of doc, panicking if not
// valid.
func FromDocument(doc io.Reader) *Schematic {
	return rawFromDocument(doc).Compact()
}

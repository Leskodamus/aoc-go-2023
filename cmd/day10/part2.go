/*
--- Part Two ---

You quickly reach the farthest point of the loop, but the animal never emerges. Maybe its nest is within the area enclosed by the loop?

To determine whether it's even worth taking the time to search for such a nest, you should calculate how many tiles are contained within the loop. For example:

...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........

The above loop encloses merely four tiles - the two pairs of . in the southwest and southeast (marked I below). The middle . tiles (marked O below) are not in the loop. Here is the same loop again with those regions marked:

...........
.S-------7.
.|F-----7|.
.||OOOOO||.
.||OOOOO||.
.|L-7OF-J|.
.|II|O|II|.
.L--JOL--J.
.....O.....

In fact, there doesn't even need to be a full tile path to the outside for tiles to count as outside the loop - squeezing between pipes is also allowed! Here, I is still within the loop and O is still outside the loop:

..........
.S------7.
.|F----7|.
.||OOOO||.
.||OOOO||.
.|L-7F-J|.
.|II||II|.
.L--JL--J.
..........

In both of the above examples, 4 tiles are enclosed by the loop.

Here's a larger example:

.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...

The above sketch has many random bits of ground, some of which are in the loop (I) and some of which are outside it (O):

OF----7F7F7F7F-7OOOO
O|F--7||||||||FJOOOO
O||OFJ||||||||L7OOOO
FJL7L7LJLJ||LJIL-7OO
L--JOL7IIILJS7F-7L7O
OOOOF-JIIF7FJ|L7L7L7
OOOOL7IF7||L7|IL7L7|
OOOOO|FJLJ|FJ|F7|OLJ
OOOOFJL-7O||O||||OOO
OOOOL---JOLJOLJLJOOO

In this larger example, 8 tiles are enclosed by the loop.

Any tile that isn't part of the main loop can count as being enclosed by the loop. Here's another example with many bits of junk pipe lying around that aren't connected to the main loop at all:

FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L

Here are just the tiles that are enclosed by the loop marked with I:

FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJIF7FJ-
L---JF-JLJIIIIFJLJJ7
|F|F-JF---7IIIL7L|7|
|FFJF7L7F-JF7IIL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L

In this last example, 10 tiles are enclosed by the loop.

Figure out whether you have time to search for the nest by calculating the area within the loop. How many tiles are enclosed by the loop?
*/

package day10

import (
	"aoc2023/internal/util"
	"fmt"
	"math"
	"slices"
	"strings"
)

func (p Part2) Run (input string) {
    buffer, err := util.ReadInput (input)
    if err != nil {
        util.ExitErr (1, err)
    }

    grid := Grid{}
    var start Point

    for y := 0; buffer.Scan (); y++ {
        line := buffer.Text()
        for x, c := range line {
            grid[Point{x, y}] = c
            if c == 'S' {
                start = Point {x, y}
            }
        }
    }

    // Get symbol for S
	grid[start] = map[[4]bool]rune{
		{true, false, true, false}: '|', {false, true, false, true}: '-',
		{true, true, false, false}: 'L', {true, false, false, true}: 'J',
		{false, false, true, true}: '7', {false, true, true, false}: 'F',
	}[[4]bool{
		strings.ContainsRune("7F|", grid[start.Add(&Point{0, -1})]),
		strings.ContainsRune("-7J", grid[start.Add(&Point{1, 0})]),
		strings.ContainsRune("JL|", grid[start.Add(&Point{0, 1})]),
		strings.ContainsRune("-FL", grid[start.Add(&Point{-1, 0})]),
	}]

    var path []Point
    var area int = 0

    for point, next := start, start; point == start || next != start; 
            path = append (path, point) {

        point, next = next, start

        for _, dir := range Directions[grid[point]] {
            if !slices.Contains (path, point.Add (&dir)) {
                next = point.Add (&dir)
            }
        }

        // https://en.wikipedia.org/wiki/Shoelace_formula
        area += point.x * next.y - point.y * next.x
    }

    // Given the area size calculated with the shoelace formula, the number of
    // tiles can be calculaed applying the Pick's theorem, by calculating i.
    // Where i = the number of integer points within the polygon.
    // Formula: A = i + b/2 - 1 => i = A - b/2 + 1
    // https://en.wikipedia.org/wiki/Pick%27s_theorem 
    area = int (math.Abs (float64 (area / 2)))
    number_of_points := area - (len (path) / 2 - 1)
    fmt.Println (number_of_points)
}


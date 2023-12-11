/*
--- Day 11: Cosmic Expansion ---

You continue following signs for "Hot Springs" and eventually come across an observatory. The Elf within turns out to be a researcher studying cosmic expansion using the giant telescope here.

He doesn't know anything about the missing machine parts; he's only visiting for this research project. However, he confirms that the hot springs are the next-closest area likely to have people; he'll even take you straight there once he's done with today's observation analysis.

Maybe you can help him with the analysis to speed things up?

The researcher has collected a bunch of data and compiled the data into a single giant image (your puzzle input). The image includes empty space (.) and galaxies (#). For example:

...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....

The researcher is trying to figure out the sum of the lengths of the shortest path between every pair of galaxies. However, there's a catch: the universe expanded in the time it took the light from those galaxies to reach the observatory.

Due to something involving gravitational effects, only some space expands. In fact, the result is that any rows or columns that contain no galaxies should all actually be twice as big.

In the above example, three columns and two rows contain no galaxies:

   v  v  v
 ...#......
 .......#..
 #.........
>..........<
 ......#...
 .#........
 .........#
>..........<
 .......#..
 #...#.....
   ^  ^  ^

These rows and columns need to be twice as big; the result of cosmic expansion therefore looks like this:

....#........
.........#...
#............
.............
.............
........#....
.#...........
............#
.............
.............
.........#...
#....#.......

Equipped with this expanded universe, the shortest path between every pair of galaxies can be found. It can help to assign every galaxy a unique number:

....1........
.........2...
3............
.............
.............
........4....
.5...........
............6
.............
.............
.........7...
8....9.......

In these 9 galaxies, there are 36 pairs. Only count each pair once; order within the pair doesn't matter. For each pair, find any shortest path between the two galaxies using only steps that move up, down, left, or right exactly one . or # at a time. (The shortest path between two galaxies is allowed to pass through another galaxy.)

For example, here is one of the shortest paths between galaxies 5 and 9:

....1........
.........2...
3............
.............
.............
........4....
.5...........
.##.........6
..##.........
...##........
....##...7...
8....9.......

This path has length 9 because it takes a minimum of nine steps to get from galaxy 5 to galaxy 9 (the eight locations marked # plus the step onto galaxy 9 itself). Here are some other example shortest path lengths:

    Between galaxy 1 and galaxy 7: 15
    Between galaxy 3 and galaxy 6: 17
    Between galaxy 8 and galaxy 9: 5

In this example, after expanding the universe, the sum of the shortest path between all 36 pairs of galaxies is 374.

Expand the universe, then find the length of the shortest path between every pair of galaxies. What is the sum of these lengths?
*/

package day11

import (
	"aoc2023/internal/util"
	"fmt"
	"math"
)


type Part1 struct { util.Part }
type Part2 struct { util.Part }

var Challenge util.Challenge = util.Challenge {
    Part1: Part1{}, Part2: Part2{},
}


type PointPair struct {
    A, B util.Point
}


func (p *Part1) expand_galaxy (grid *[][]int) {
    for i := 0; i < len (*grid); i++ {
        if p.has_no_galaxies ((*grid)[i]) {
            *grid = append ((*grid)[:i+1], (*grid)[i:]...)
            (*grid)[i] = make ([]int, len ((*grid)[i+1]))
            i++
        }
    }

    for i := 0; i < len ((*grid)[0]); i++ {
        col := []int {}
        for j := 0; j < len (*grid); j++ {
            col = append (col, (*grid)[j][i])
        }
        if p.has_no_galaxies (col) {
            for j := 0; j < len (*grid); j++ {
                (*grid)[j] = append ((*grid)[j][:i+1], (*grid)[j][i:]...)
                (*grid)[j][i] = 0
            }
            i++
        }
    }
}


func (p *Part1) has_no_galaxies (arr []int) bool {
    for _, v := range arr {
        if v > 0 { return false }
    }
    return true
}


func (p Part1) Run (input string) {
    sum := 0

    buffer, err := util.ReadInput (input)
    if err != nil {
        util.ExitErr (1, err)
    }

    grid := [][]int {}
    ngalaxies := 1

    for y := 0; buffer.Scan (); y++ {
        line := buffer.Text ()
        grid = append (grid, []int {})

        for _, c := range line {
            if c == '#' {
                grid[y] = append (grid[y], ngalaxies)
                ngalaxies++
            } else {
                grid[y] = append (grid[y], 0)
            }
        }
    }

    p.expand_galaxy (&grid)

    // Use vertices instead
    galaxies := []util.Point{}

    for y := 0; y < len (grid); y++ {
        for x := 0; x < len (grid[y]); x++ {
            if grid[y][x] > 0 {
                galaxies = append (galaxies, util.Point {x, y})
            }
        }
    }

    // Calculate distances for each point to point pair
    distances := map[PointPair]int{}

    for i := 0; i < len (galaxies); i++ {
        for j := 0; j < len (galaxies); j++ {
            if i == j { continue }
            dx := int (math.Abs (float64 (galaxies[i].X - galaxies[j].X)))
            dy := int (math.Abs (float64 (galaxies[i].Y - galaxies[j].Y)))
            dist := dx + dy

            pair := PointPair {galaxies[i], galaxies[j]}
            rev_pair := PointPair {galaxies[j], galaxies[i]}

            _, has_pair := distances[pair]
            _, has_rev_pair := distances[rev_pair]

            if !has_pair && !has_rev_pair {
                distances[pair] = dist
                sum += dist
            }
        }
    }

    fmt.Println ("Sum:", sum)
}


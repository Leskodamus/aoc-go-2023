/*
--- Part Two ---

The galaxies are much older (and thus much farther apart) than the researcher initially estimated.

Now, instead of the expansion you did before, make each empty row or column one million times larger. That is, each empty row should be replaced with 1000000 empty rows, and each empty column should be replaced with 1000000 empty columns.

(In the example above, if each empty row or column were merely 10 times larger, the sum of the shortest paths between every pair of galaxies would be 1030. If each empty row or column were merely 100 times larger, the sum of the shortest paths between every pair of galaxies would be 8410. However, your universe will need to expand far beyond these values.)

Starting with the same initial image, expand the universe according to these new rules, then find the length of the shortest path between every pair of galaxies. What is the sum of these lengths?
*/

package day11

import (
	"aoc2023/internal/util"
	"fmt"
	"math"
)


var part1 Part1


func (p *Part2) has_no_galaxies_col (grid *[][]int, x int) bool {
    for y := 0; y < len (*grid); y++ {
        if (*grid)[y][x] > 0 {
            return false
        }
    }
    return true
}


func (p Part2) Run (input string) {
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
            
            // Check how many rows and columns are between the two points
            point_a := galaxies[i]
            point_b := galaxies[j]

            nrows := 0
            if point_a.Y > point_b.Y {
                // starting from point_b going to point_a
                for y := point_a.Y; y > point_b.Y; y-- {
                    if part1.has_no_galaxies (grid[y]) {
                        nrows += 1000000 - 1
                    }
                }
            } else {
                // starting from point_a going to point_b
                for y := point_a.Y; y < point_b.Y; y++ {
                    if part1.has_no_galaxies (grid[y]) {
                        nrows += 1000000 - 1}
                }
            }

            ncols := 0
            if point_a.X > point_b.X {
                // starting from point_b going to point_a
                for x := point_a.X; x > point_b.X; x-- {
                    if p.has_no_galaxies_col (&grid, x) {
                        ncols += 1000000 - 1
                    }
                }
            } else {
                // starting from point_a going to point_b
                for x := point_a.X; x < point_b.X; x++ {
                    if p.has_no_galaxies_col (&grid, x) {
                        ncols += 1000000 - 1
                    }
                }
            }

            dx := int (math.Abs (float64 (galaxies[i].X - galaxies[j].X)))
            dy := int (math.Abs (float64 (galaxies[i].Y - galaxies[j].Y)))
            dist := dx + dy + ncols + nrows

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


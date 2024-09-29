/*
--- Part Two ---
The parabolic reflector dish deforms, but not in a way that focuses the beam. To do that, you'll need to move the rocks to the edges of the platform. Fortunately, a button on the side of the control panel labeled "spin cycle" attempts to do just that!

Each cycle tilts the platform four times so that the rounded rocks roll north, then west, then south, then east. After each tilt, the rounded rocks roll as far as they can before the platform tilts in the next direction. After one cycle, the platform will have finished rolling the rounded rocks in those four directions in that order.

Here's what happens in the example above after each of the first few cycles:

After 1 cycle:
.....#....
....#...O#
...OO##...
.OO#......
.....OOO#.
.O#...O#.#
....O#....
......OOOO
#...O###..
#..OO#....

After 2 cycles:
.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#..OO###..
#.OOO#...O

After 3 cycles:
.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#...O###.O
#.OOO#...O
This process should work if you leave it running long enough, but you're still worried about the north support beams. To make sure they'll survive for a while, you need to calculate the total load on the north support beams after 1000000000 cycles.

In the above example, after 1000000000 cycles, the total load on the north support beams is 64.

Run the spin cycle for 1000000000 cycles. Afterward, what is the total load on the north support beams?
*/

package day14

import (
	"aoc2023/internal/util"
	"fmt"
	"strings"
)

/*
 * Move all rounded rocks to the north.
*/
func (p *Part2) move_rounded_rocks_north (field [][]rune) {
    for r := 0; r < len (field); r++ {
        for c := 0; c < len (field[r]); c++ {
            // Move any 'O' to the top of the field
            if field[r][c] == 'O' {
                for i := r; i > 0; i-- {
                    if field[i-1][c] == '#' || field[i-1][c] == 'O' {
                        break
                    }
                    field[i][c] = '.'
                    field[i-1][c] = 'O'
                }
            }
        }
    }
}

// Move all rounded rocks to the west.
func (p *Part2) move_rounded_rocks_west (field [][]rune) {
	for r := 0; r < len (field); r++ {
		for c := 0; c < len (field[r]); c++ {
			// Move any 'O' to the left of the field
			if field[r][c] == 'O' {
				for i := c; i > 0; i-- {
					if field[r][i-1] == '#' || field[r][i-1] == 'O' {
						break
					}
					field[r][i] = '.'
					field[r][i-1] = 'O'
				}
			}
		}
	}
}

// Move all rounded rocks to the south.
func (p *Part2) move_rounded_rocks_south (field [][]rune) {
	for r := len (field) - 1; r >= 0; r-- {
		for c := 0; c < len (field[r]); c++ {
			// Move any 'O' to the bottom of the field
			if field[r][c] == 'O' {
				for i := r; i < len (field) - 1; i++ {
					if field[i+1][c] == '#' || field[i+1][c] == 'O' {
						break
					}
					field[i][c] = '.'
					field[i+1][c] = 'O'
				}
			}
		}
	}
}

// Move all rounded rocks to the east.
func (p *Part2) move_rounded_rocks_east (field [][]rune) {
	for r := 0; r < len (field); r++ {
		for c := len (field[r]) - 1; c >= 0; c-- {
			// Move any 'O' to the right of the field
			if field[r][c] == 'O' {
				for i := c; i < len (field[r]) - 1; i++ {
					if field[r][i+1] == '#' || field[r][i+1] == 'O' {
						break
					}
					field[r][i] = '.'
					field[r][i+1] = 'O'
				}
			}
		}
	}
}

func (p *Part2) cycle(field [][]rune) {
	p.move_rounded_rocks_north(field)
	p.move_rounded_rocks_west(field)
	p.move_rounded_rocks_south(field)
	p.move_rounded_rocks_east(field)
}

type seenField struct {
	index int
	field [][]rune
}

func (p Part2) Run (input string) {
    sum := 0

    buffer, err := util.ReadInput (input)
    if err != nil {
        util.ExitErr (1, err)
    }

    field := make([][]rune, 0)

    for buffer.Scan() {
        line := buffer.Text()
        field = append (field, []rune(line)) 
    }

	const cycles = 1000000000

	seenIndexes := make(map[string]int)
	seenFields := make(map[int][][]rune)

	i, start := 0, 0
	for ; i < cycles; i++ {
		p.cycle(field)

		k := util.RunesOfRunesSliceToString(field)
		if idx, ok := seenIndexes[k]; ok {
			start = idx
			i++
			break
		}

		seenIndexes[k] = i + 1
		seenFields[i+1] = make([][]rune, len(field))
		// Copy the field (to avoid reference)
		for j, row := range field {
			seenFields[i+1][j] = make([]rune, len(row))
			copy(seenFields[i+1][j], row)
		}
	}

	period := i - start
	field = seenFields[start + ((cycles - i - 1) % period) + 1]

    for i, row := range field {
        // Count number of rounded rocks in the row
        n_o_rocks := strings.Count (string(row), "O")
        sum += n_o_rocks * (len(field) - i)
    }

    fmt.Println ("Sum:", sum)
}


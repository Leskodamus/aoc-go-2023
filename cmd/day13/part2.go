/*
--- Part Two ---

You resume walking through the valley of mirrors and - SMACK! - run directly into one. Hopefully nobody was watching, because that must have been pretty embarrassing.

Upon closer inspection, you discover that every mirror has exactly one smudge: exactly one . or # should be the opposite type.

In each pattern, you'll need to locate and fix the smudge that causes a different reflection line to be valid. (The old reflection line won't necessarily continue being valid after the smudge is fixed.)

Here's the above example again:

#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#

The first pattern's smudge is in the top-left corner. If the top-left # were instead ., it would have a different, horizontal line of reflection:

1 ..##..##. 1
2 ..#.##.#. 2
3v##......#v3
4^##......#^4
5 ..#.##.#. 5
6 ..##..##. 6
7 #.#.##.#. 7

With the smudge in the top-left corner repaired, a new horizontal line of reflection between rows 3 and 4 now exists. Row 7 has no corresponding reflected row and can be ignored, but every other row matches exactly: row 1 matches row 6, row 2 matches row 5, and row 3 matches row 4.

In the second pattern, the smudge can be fixed by changing the fifth symbol on row 2 from . to #:

1v#...##..#v1
2^#...##..#^2
3 ..##..### 3
4 #####.##. 4
5 #####.##. 5
6 ..##..### 6
7 #....#..# 7

Now, the pattern has a different horizontal line of reflection between rows 1 and 2.

Summarize your notes as before, but instead use the new different reflection lines. In this example, the first pattern's new horizontal line has 3 rows above it and the second pattern's new horizontal line has 1 row above it, summarizing to the value 400.

In each pattern, fix the smudge and find the different line of reflection. What number do you get after summarizing the new reflection line in each pattern in your notes?
*/

package day13

import (
    "aoc2023/internal/util"
    "fmt"
)


var part1 Part1 = Part1{}


/*
 * Returns the number of differences in the reflection line or 0 if there is none.
*/
func (p *Part2) find_one_off_reflection (pattern []string) int {
    width := len (pattern[0])

    for col := 0; col < width - 1; col++ {
        ndiff := 0

        for _, line := range pattern {
            // Compare left and right reflection
            for i, j := col, col+1; i >= 0 && j < width; i, j = i-1, j+1 {
                left, right := line[i], line[j]
                if left != right {
                    ndiff++
                }
            }
            if ndiff > 1 { break }
        }
        if ndiff == 1 { return col+1 }
    }

    return 0
}


func (p Part2) Run (input string) {
    sum := 0

    buffer, err := util.ReadInput (input)
    if err != nil {
        util.ExitErr (1, err)
    }

    block := []string{}
    field := [][]string{}

    for buffer.Scan() {
        line := buffer.Text()
        if len (line) == 0 {
            field = append (field, block)
            block = []string{}
            continue
        }
        block = append (block, line)
    }

    // Check if there is a block left. 
    // Happens if the input file does not end with a newline.
    if len (block) > 0 {
        field = append (field, block)
        block = []string{}
    }

    for _, pattern := range field {
        sum += p.find_one_off_reflection (pattern)

        tposed := part1.transpose (pattern)
        sum += p.find_one_off_reflection (tposed) * 100
    }

    fmt.Println ("Sum:", sum)
}


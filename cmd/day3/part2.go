/*
--- Part Two ---

The engineer finds the missing part and installs it in the engine! As the engine springs to life, you jump in the closest gondola, finally ready to ascend to the water source.

You don't seem to be going very fast, though. Maybe something is still wrong? Fortunately, the gondola has a phone labeled "help", so you pick it up and the engineer answers.

Before you can explain the situation, she suggests that you look out the window. There stands the engineer, holding a phone in one hand and waving with the other. You're going so slowly that you haven't even left the station. You exit the gondola.

The missing part wasn't the only issue - one of the gears in the engine is wrong. A gear is any * symbol that is adjacent to exactly two part numbers. Its gear ratio is the result of multiplying those two numbers together.

This time, you need to find the gear ratio of every gear and add them all up so that the engineer can figure out which gear needs to be replaced.

Consider the same engine schematic again:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..

In this schematic, there are two gears. The first is in the top left; it has part numbers 467 and 35, so its gear ratio is 16345. The second gear is in the lower right; its gear ratio is 451490. (The * adjacent to 617 is not a gear because it is only adjacent to one part number.) Adding up all of the gear ratios produces 467835.

What is the sum of all of the gear ratios in your engine schematic?
*/

package day3

import (
	"aoc2023/internal/util"
	"fmt"
	"unicode"
)


func (p *Part2) get_number (row []rune, i int) int {
    var number int = 0

    // Check digits to left
    var incr int = 1
    for k := i; k >= 0 && unicode.IsDigit (row[k]); k-- {
        digit := int (row[k] - '0')
        number = digit * incr + number
        incr *= 10
    }

    // Check digits to right
    for k := i+1; k < len (row) && unicode.IsDigit (row[k]); k++ {
        digit := int (row[k] - '0')
        number = number * 10 + digit
    }

    return number
}


func (p Part2) Run (input string) {
    sum := 0

    buffer, err := util.ReadInput (input)
    if err != nil {
        util.ExitErr (1, err)
    }

    var schematic = [][]rune{}

    for buffer.Scan() {
        schematic = append (schematic, []rune (buffer.Text()))
    }

    // Get * symbols which are adjacent to exactly two numbers
    for i := 0; i < len (schematic); i++ {
        for j := 0; j < len (schematic[i]); j++ {
            var nAdjacent int = 0
            var numbers []int

            if schematic[i][j] == '*' {
                // Check if digit is to the left
                if j > 1 && unicode.IsDigit (schematic[i][j-1]) {
                    nAdjacent++
                    numbers = append (numbers, p.get_number (schematic[i], j-1))
                }

                // Check if digit is to the right
                if j < len (schematic[i])+1 && unicode.IsDigit (schematic[i][j+1]) {
                    nAdjacent++
                    numbers = append (numbers, p.get_number (schematic[i], j+1))
                } 
                
                // Check if digit is above
                if i > 0 {
                    if unicode.IsDigit (schematic[i-1][j]) {
                        nAdjacent++
                        numbers = append (numbers, p.get_number (schematic[i-1], j))
                    } else {
                        // Only need to check left and right if there is no digit 
                        // directly above the asterisk

                        // Check diagonals above left
                        if j > 0 && unicode.IsDigit (schematic[i-1][j-1]) {
                            nAdjacent++
                            numbers = append (numbers, p.get_number (schematic[i-1], j-1))
                        } 

                        // Check diagonals above right
                        if j < len (schematic[i-1]) - 1 && 
                                unicode.IsDigit (schematic[i-1][j+1]) {
                            nAdjacent++
                            numbers = append (numbers, p.get_number (schematic[i-1], j+1))
                        }
                    }
                }

                // Check if digit is below
                if i < len (schematic) - 1 {
                    if unicode.IsDigit (schematic[i+1][j]) {
                        nAdjacent++
                        numbers = append (numbers, p.get_number (schematic[i+1], j))
                    } else {
                        // Check diagonals below left
                        if j > 0 && unicode.IsDigit (schematic[i+1][j-1]) {
                            nAdjacent++
                            numbers = append (numbers, p.get_number (schematic[i+1], j-1))
                        }

                        // Check diagonals below right
                        if j < len (schematic[i+1]) - 1 && 
                                unicode.IsDigit (schematic[i+1][j+1]) {
                            nAdjacent++
                            numbers = append (numbers, p.get_number (schematic[i+1], j+1))
                        }
                    }
                }
            }

            if nAdjacent == 2 {
                sum += numbers[0] * numbers[1]
            }
        }
    }
   
    fmt.Println ("Sum:", sum)
}



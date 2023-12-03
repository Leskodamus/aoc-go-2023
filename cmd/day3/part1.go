/*
--- Day 3: Gear Ratios ---

You and the Elf eventually reach a gondola lift station; he says the gondola lift will take you up to the water source, but this is as far as he can bring you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem: they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of surprise. "Sorry, I wasn't expecting anyone! The gondola lift isn't working right now; it'll still be a while before I can fix it." You offer to help.

The engineer explains that an engine part seems to be missing from the engine, but nobody can figure out which one. If you can add up all the part numbers in the engine schematic, it should be easy to work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of the engine. There are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a symbol, even diagonally, is a "part number" and should be included in your sum. (Periods (.) do not count as a symbol.)

Here is an example engine schematic:

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

In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58 (middle right). Every other number is adjacent to a symbol and so is a part number; their sum is 4361.

Of course, the actual engine schematic is much larger. What is the sum of all of the part numbers in the engine schematic?
*/

package day3

import (
	"aoc2023/internal/util"
	"fmt"
	"unicode"
)


type Part1 struct { util.Part }
type Part2 struct { util.Part }

var Day3 util.ChallengeDay = util.ChallengeDay {
    Part1: Part1{}, Part2: Part2{},
}


func (p *Part1) has_adjacent_symbol () bool {
    return true 
}


func (p Part1) Run (input string) {
    sum := 0

    buffer, err := util.ReadInput (input)
    if err != nil {
        util.ExitErr (1, err)
    }

    var schematic = [][]rune{}

    for buffer.Scan() {
        schematic = append (schematic, []rune (buffer.Text()))
    }

    // Get numbers which do not have a symbol adjacent to them
    for i := 0; i < len (schematic); i++ {
        for j := 0; j < len (schematic[i]); j++ {
            var number int = 0
            var nDigits int = 0

            if unicode.IsDigit (schematic[i][j]) {
                for j < len (schematic[i]) && unicode.IsDigit (schematic[i][j]) {
                    number = number * 10 + int (schematic[i][j] - '0')
                    nDigits++
                    j++
                }
            
                begNum := j - nDigits
                endNum := j - 1

                // Check to the left
                if j > 1 && begNum > 0 && schematic[i][begNum-1] != '.' {
                    sum += number
                    continue
                }

                // Check to the right
                if j < len (schematic[i]) && schematic[i][endNum+1] != '.' {
                    sum += number
                    continue
                } 
                
                // Check above
                if i > 0 {
                    toContinue := false
                    for k := begNum; k <= endNum; k++ {
                        if schematic[i-1][k] != '.' && 
                                !unicode.IsDigit (schematic[i-1][k]) {
                            sum += number
                            toContinue = true
                        }
                    }
                    if toContinue { continue }

                    // Check diagonals above left
                    if begNum > 0 && schematic[i-1][begNum-1] != '.' &&
                            !unicode.IsDigit (schematic[i-1][begNum-1]) {
                        sum += number
                        continue
                    }

                    // Check diagonals above right
                    if endNum < len (schematic[i-1]) - 1 &&
                            schematic[i-1][endNum+1] != '.' &&
                            !unicode.IsDigit (schematic[i-1][endNum+1]) {
                        sum += number
                        continue
                    }
                }

                // Check below
                if i < len (schematic) - 1 {
                    toContinue := false
                    for k := begNum; k <= endNum; k++ {
                        if schematic[i+1][k] != '.' && 
                                !unicode.IsDigit (schematic[i+1][k]) {
                            sum += number
                            toContinue = true
                        }
                    }
                    if toContinue { continue }

                    if begNum > 0 && schematic[i+1][begNum-1] != '.' &&
                            !unicode.IsDigit (schematic[i+1][begNum-1]) {
                        sum += number
                        continue
                    }

                    if endNum < len (schematic[i+1]) - 1 &&
                            schematic[i+1][endNum+1] != '.' &&
                            !unicode.IsDigit (schematic[i+1][endNum+1]) {
                        sum += number
                        continue
                    }
                }
            }
        }
    }
   
    fmt.Println ("Sum:", sum)
}


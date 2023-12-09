/*
--- Part Two ---

Of course, it would be nice to have even more history included in your report. Surely it's safe to just extrapolate backwards as well, right?

For each history, repeat the process of finding differences until the sequence of differences is entirely zero. Then, rather than adding a zero to the end and filling in the next values of each previous sequence, you should instead add a zero to the beginning of your sequence of zeroes, then fill in new first values for each previous sequence.

In particular, here is what the third example history looks like when extrapolating back in time:

5  10  13  16  21  30  45
  5   3   3   5   9  15
   -2   0   2   4   6
      2   2   2   2
        0   0   0

Adding the new values on the left side of each sequence from bottom to top eventually reveals the new left-most history value: 5.

Doing this for the remaining example data above results in previous values of -3 for the first history and 0 for the second history. Adding all three new values together produces 2.

Analyze your OASIS report again, this time extrapolating the previous value for each history. What is the sum of these extrapolated values?
*/

package day9

import (
	"aoc2023/internal/util"
	"fmt"
	"strings"
)


var part1 Part1


func (p Part2) Run (input string) {
    sum := 0
    buffer, err := util.ReadInput (input)
    if err != nil {
        util.ExitErr (1, err)
    }

    var initial_signals [][]int
    var extrapolate_signals [][][]int

    for buffer.Scan () {
        line := buffer.Text ()
        history := util.StrFieldsToInts (strings.Fields (line))
        initial_signals = append (initial_signals, history)
    }

    for _, initial := range initial_signals {
        sequences := [][]int { initial }
        
        for i := 0; i < len (sequences) ; i++ {
            // Generate sequence for initial signal
            var sequence []int

            for j := 1; j < len (sequences[i]); j++ {
                diff := sequences[i][j] - sequences[i][j-1]
                sequence = append (sequence, diff)
            }

            sequences = append (sequences, sequence)

            if part1.arr_is_all_zeroes (sequence) {
                sequences[i+1] = append ([]int {0}, sequences[i+1]...)
                break
            }
        }
        extrapolate_signals = append (extrapolate_signals, sequences)
    }

    // For each list of sequences
    for s := 0; s < len (extrapolate_signals); s++ {
        // For each sequence in list, starting from bottom
        for i := len (extrapolate_signals[s]) - 1; i > 0; i-- {
            value := extrapolate_signals[s][i-1][0] -
                    extrapolate_signals[s][i][0]

            // Insert calculated value at the beginning of the sequence
            extrapolate_signals[s][i-1] = append ([]int{value}, extrapolate_signals[s][i-1]...)
        }
    }

    // Calculate sum
    for _, extrapolate := range extrapolate_signals {
        sum += extrapolate[0][0]
    }

    fmt.Println ("Sum:", sum)
}


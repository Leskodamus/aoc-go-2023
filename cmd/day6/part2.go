/*
--- Part Two ---

As the race is about to start, you realize the piece of paper with race times and record distances you got earlier actually just has very bad kerning. There's really only one race - ignore the spaces between the numbers on each line.

So, the example from before:

Time:      7  15   30
Distance:  9  40  200

...now instead means this:

Time:      71530
Distance:  940200

Now, you have to figure out how many ways there are to win this single race. In this example, the race lasts for 71530 milliseconds and the record distance you need to beat is 940200 millimeters. You could hold the button anywhere from 14 to 71516 milliseconds and beat the record, a total of 71503 ways!

How many ways can you beat the record in this one much longer race?
*/

package day6

import (
	"aoc2023/internal/util"
	"fmt"
	"strconv"
	"strings"
)


func (p Part2) Run (input string) {
    sum := 1

    buffer, err := util.ReadInput (input)
    if err != nil {
        util.ExitErr (1, err)
    }

    buffer.Scan()
    times_line := buffer.Text()

    buffer.Scan()
    distances_line := buffer.Text()

    record_time, _ := strconv.Atoi (strings.Replace (strings.TrimPrefix (times_line, "Time: "), " ", "", -1))
    record_distance, _ := strconv.Atoi (strings.Replace (strings.TrimPrefix (distances_line, "Distance: "), " ", "", -1))

    nways := 0
    for j := 0; j < record_time; j++ {
        distance := j * (record_time - j)
        if distance > record_distance { nways += 1 }
    }
    sum *= nways

    fmt.Println ("Sum: ", sum)
}


/*
--- Part Two ---

As you look out at the field of springs, you feel like there are way more springs than the condition records list. When you examine the records, you discover that they were actually folded up this whole time!

To unfold the records, on each row, replace the list of spring conditions with five copies of itself (separated by ?) and replace the list of contiguous groups of damaged springs with five copies of itself (separated by ,).

So, this row:

.# 1

Would become:

.#?.#?.#?.#?.# 1,1,1,1,1

The first line of the above example would become:

???.###????.###????.###????.###????.### 1,1,3,1,1,3,1,1,3,1,1,3,1,1,3

In the above example, after unfolding, the number of possible arrangements for some rows is now much larger:

    ???.### 1,1,3 - 1 arrangement
    .??..??...?##. 1,1,3 - 16384 arrangements
    ?#?#?#?#?#?#?#? 1,3,1,6 - 1 arrangement
    ????.#...#... 4,1,1 - 16 arrangements
    ????.######..#####. 1,6,5 - 2500 arrangements
    ?###???????? 3,2,1 - 506250 arrangements

After unfolding, adding all of the possible arrangement counts together produces 525152.

Unfold your condition records; what is the new sum of possible arrangement counts?
*/

package day12

import (
	"aoc2023/internal/util"
	"fmt"
	"regexp"
	"strings"
)


var part1 Part1


func (p Part2) Run (input string) {
    sum := 0

    buffer, err := util.ReadInput (input)
    if err != nil {
        util.ExitErr (1, err)
    }

    records := []Record{}

    for buffer.Scan () {
        line := buffer.Text ()
        parts := strings.Split (line, " ")
        records = append (records, Record {
            conditions: parts[0],
            group: util.StrFieldsToInts (strings.Split (parts[1], ",")),
        })
    }

    for _, record := range records {
        conditions := record.conditions
        group := record.group
        for i := 0; i < 4; i++ {
            record.conditions = fmt.Sprintf ("%s?%s", record.conditions, conditions)
            record.group = append (record.group, group...)
        }

        re := []*regexp.Regexp{}
        for _, g := range record.group {
            re_str := fmt.Sprintf ("[.?][#?]{%d}[.?]", g)
            re = append (re, regexp.MustCompile (re_str))
        }

        sum += part1.count_arrangements (fmt.Sprintf (".%s.", record.conditions), re)
    }

    fmt.Println (sum)
}


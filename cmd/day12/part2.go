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
	"slices"
	"strings"
)


var part1 Part1


func (p *Part2) is_valid (conditions string, groups []int) bool {
    condition_parts := strings.Split (conditions, ".")
    count_spring_parts := []int{}

    for _, part := range condition_parts {
        c := strings.Count (part, "#")
        if c == 0 { continue }
        count_spring_parts = append (count_spring_parts, c)
    }

    return slices.Equal (count_spring_parts, groups)
}


func (p *Part2) count_arrangements (conditions string, groups []int) int {
    if !strings.Contains (conditions, "?") {
        if p.is_valid (conditions, groups) {
            return 1
        }
        return 0
    }

    i := strings.Index (conditions, "?")
    count := 0
    count += p.count_arrangements (conditions[:i] + "#" + conditions[i+1:], groups)
    count += p.count_arrangements (conditions[:i] + "." + conditions[i+1:], groups)

    return count
}


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
        sum += p.count_arrangements (conditions, group)
    }

    fmt.Println (sum)
}


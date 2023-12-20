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
	"strings"
)


type cacheKey struct {
    conditions string
    groups     string   // This is a string because slices are not comparable.
    groupLoc   int
}

var cache = make (map[cacheKey]int)


func (p *Part2) count_arrangements (conditions string, groups []int, group_loc int) int {
    key := cacheKey{ conditions, util.IntSlicetoString(groups, ","), group_loc }

    if val, ok := cache[key]; ok {
        return val
    }

    if len (conditions) == 0 {
        if len (groups) == 0 && group_loc == 0 {
            cache[key] = 1
            return 1
        }
        cache[key] = 0
        return 0
    }

    count := 0

    possibilities := []rune{rune (conditions[0])}
    if conditions[0] == '?' {
        possibilities = []rune{ '.', '#' }
    }

    for _, pos := range possibilities {
        if pos == '#' {
            count += p.count_arrangements (conditions[1:], groups, group_loc + 1)
        } else {
            if group_loc > 0 {
                if len (groups) > 0 && groups[0] == group_loc {
                    count += p.count_arrangements (conditions[1:], groups[1:], 0)
                }
            } else {
                count += p.count_arrangements (conditions[1:], groups, 0)
            }
        }
    }

    cache[key] = count

    return count
}


func (p Part2) Run (input string) {
    sum := 0

    buffer, err := util.ReadInput (input)
    if err != nil {
        util.ExitErr (1, err)
    }

    for buffer.Scan () {
        line := buffer.Text ()
        parts := strings.Split (line, " ")
        
        conditions, group := func() (string, []int) {
            original_parts := parts[0]
            original_groups := parts[1]

            for i := 0; i < 4; i++ {
                parts[0] += "?" + original_parts
                parts[1] += "," + original_groups
            }
            // Dot is appended to the conditions to allow iterating up to the last symbol.
            // This symbol could be anything, even just a space.
            return parts[0] + ".", util.StrFieldsToInts(strings.Split (parts[1], ","))
        }()

        sum += p.count_arrangements (conditions, group, 0)
    }

    fmt.Println (sum)
}


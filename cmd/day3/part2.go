/*
*/

package day3

import (
    "aoc2023/internal/util"
)


func (p Part2) Run (input string) {
    buffer, err := util.ReadInput (input)
    if err != nil {
        util.ExitErr (1, err)
    }

    for buffer.Scan() {
    }
}


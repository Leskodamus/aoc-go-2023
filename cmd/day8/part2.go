/*
--- Part Two ---

The sandstorm is upon you and you aren't any closer to escaping the wasteland. You had the camel follow the instructions, but you've barely left your starting position. It's going to take significantly more steps to escape!

What if the map isn't for people - what if the map is for ghosts? Are ghosts even bound by the laws of spacetime? Only one way to find out.

After examining the maps a bit longer, your attention is drawn to a curious fact: the number of nodes with names ending in A is equal to the number ending in Z! If you were a ghost, you'd probably just start at every node that ends with A and follow all of the paths at the same time until they all simultaneously end up at nodes that end with Z.

For example:

LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)

Here, there are two starting nodes, 11A and 22A (because they both end with A). As you follow each left/right instruction, use that instruction to simultaneously navigate away from both nodes you're currently on. Repeat this process until all of the nodes you're currently on end with Z. (If only some of the nodes you're on end with Z, they act like any other node and you continue as normal.) In this example, you would proceed as follows:

    Step 0: You are at 11A and 22A.
    Step 1: You choose all of the left paths, leading you to 11B and 22B.
    Step 2: You choose all of the right paths, leading you to 11Z and 22C.
    Step 3: You choose all of the left paths, leading you to 11B and 22Z.
    Step 4: You choose all of the right paths, leading you to 11Z and 22B.
    Step 5: You choose all of the left paths, leading you to 11B and 22C.
    Step 6: You choose all of the right paths, leading you to 11Z and 22Z.

So, in this example, you end up entirely on nodes that end in Z after 6 steps.

Simultaneously start on every node that ends with A. How many steps does it take before you're only on nodes that end with Z?
*/

package day8

import (
	"aoc2023/internal/util"
	"fmt"
	"strings"
)


func (p Part2) gcd (a, b int) int {
    if b == 0 {
        return a
    }
    return p.gcd (b, a % b)
}


func (p Part2) Run (input string) {
    buffer, err := util.ReadInput (input)
    if err != nil {
        util.ExitErr (1, err)
    }

    var instructions string

    buffer.Scan()
    instructions = buffer.Text()

    buffer.Scan() // empty line

    nodes := make (map[string][2]string)
    var starting_nodes []string

    for buffer.Scan() {
        line := buffer.Text()
        if len (line) == 0 { continue }

        name := line[0:3]
        left := line[7:10]
        right := line[12:15]

        if name[2] == 'A' { 
            starting_nodes = append (starting_nodes, name) 
        }

        nodes[name] = [2]string{left, right}
    }

    var cycles []int

    for _, node := range starting_nodes {
        var i, cycle int = 0, 0
        var first_z string

        for {
            for !strings.HasSuffix (node, "Z") {
                if instructions[i] == 'L' {
                    node = nodes[node][0]
                } else if (instructions[i] == 'R') {
                    node = nodes[node][1]
                }

                i = (i+1) % len (instructions)
                cycle++
            }

            if first_z == "" {
                first_z = node
            } else if node == first_z { break } 
        }

        cycles = append(cycles, cycle)
    }

    steps := cycles[0]
    for i := 0; i < len (cycles); i++ {
        steps = (steps * cycles[i]) / p.gcd (steps, cycles[i])
    }

    fmt.Println (steps)
}


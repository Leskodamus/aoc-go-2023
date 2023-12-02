/*
--- Part Two ---

The Elf says they've stopped producing snow because they aren't getting any water! He isn't sure why the water stopped; however, he can show you how to get to the water source to check it out for yourself. It's just up ahead!

As you continue your walk, the Elf poses a second question: in each game you played, what is the fewest number of cubes of each color that could have been in the bag to make the game possible?

Again consider the example games from earlier:

Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

    In game 1, the game could have been played with as few as 4 red, 2 green, and 6 blue cubes. If any color had even one fewer cube, the game would have been impossible.
    Game 2 could have been played with a minimum of 1 red, 3 green, and 4 blue cubes.
    Game 3 must have been played with at least 20 red, 13 green, and 6 blue cubes.
    Game 4 required at least 14 red, 3 green, and 15 blue cubes.
    Game 5 needed no fewer than 6 red, 3 green, and 2 blue cubes in the bag.

The power of a set of cubes is equal to the numbers of red, green, and blue cubes multiplied together. The power of the minimum set of cubes in game 1 is 48. In games 2-5 it was 12, 1560, 630, and 36, respectively. Adding up these five powers produces the sum 2286.

For each game, find the minimum set of cubes that must have been present. What is the sum of the power of these sets?
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


type CubeColors struct { 
    red, green, blue int 
}


func (m CubeColors) Init() CubeColors {
    m.red = 1
    m.green = 1
    m.blue = 1
    return m
}


func (m *CubeColors) SetMax(color string, count int) {
    switch color {
        case "red": 
            m.red = max (m.red, count)
        case "green": 
            m.green = max (m.green, count)
        case "blue": 
            m.blue = max (m.blue, count)
        default: 
            fmt.Println ("Unknown color:", color)
    }
}


func (m *CubeColors) Power() int {
    return m.red * m.green * m.blue
}


func get_power_of_sets (line string) int {
    var id int = 0
    colors := CubeColors{}.Init()

    if line == "" { return 0 }

    // Get game id
    _, err := fmt.Sscanf (line, "Game %d:", &id)
    if err != nil {
        fmt.Println ("Error parsing game id from line:", line, err)
        return 0
    }

    line = strings.TrimPrefix (line, fmt.Sprintf ("Game %d: ", id))
    
    // For each set of cubes
    for _, cubes_set := range strings.Split (line, ";") {

        // Get color and count: i color, i color,...
        for _, color_count := range strings.Split (cubes_set, ",") {
            var color string
            var count int

            // Get color and count
            _, err := fmt.Sscanf (color_count, "%d %s", &count, &color)
            if err != nil {
                fmt.Println ("Error parsing color and count from the set of cubes:", 
                        color_count, err)
                return 0
            }
            
            colors.SetMax (color, count)
        }
    }

    return colors.Power()
}


func main() {
    if len (os.Args) != 2 {
        fmt.Fprintln (os.Stderr, "Usage: ", os.Args[0]," <input file>")
        os.Exit (1)
    }

    file, err := os.Open (os.Args[1])
    if err != nil {
        fmt.Println("Error reading input:", err)
        os.Exit (1)
    }
    defer file.Close()

    sum := 0
    input := bufio.NewScanner (file)

    for input.Scan() {
        sum += get_power_of_sets (input.Text())
    }

    fmt.Println ("Sum: ", sum)
}


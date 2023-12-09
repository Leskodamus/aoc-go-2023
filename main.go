package main

/*
CLI syntax:
    go run main.go <day> <part> [input file]

    <day> - day of the challenge {1..25}
    <part> - part of the challenge {1,2}
    [input file] - optional input file, defaults to input.txt in directory of <day>
*/

import (
	"flag"
	"fmt"
	"os"
    "aoc2023/internal/util"
    "aoc2023/cmd/day1"
    "aoc2023/cmd/day2"
    "aoc2023/cmd/day3"
    "aoc2023/cmd/day4"
    "aoc2023/cmd/day5"
    "aoc2023/cmd/day6"
    "aoc2023/cmd/day7"
    "aoc2023/cmd/day8"
    "aoc2023/cmd/day9"
)


func main() {
    flag.Usage = func () {
        fmt.Fprintf (flag.CommandLine.Output(), "Usage: %s <day> <part> [input file]\n", 
                os.Args[0])
        flag.PrintDefaults()
    }

    // Parse CLI arguments
    var day, part int
    var input string

    flag.IntVar (&day, "day", 0, "day of the challenge {1..25}")
    flag.IntVar (&part, "part", 0, "part of the challenge {1,2}")
    flag.StringVar (&input, "input", "", "input txt file")

    flag.Parse()

    if flag.NArg() == 0 && day == 0 && part == 0 {
        flag.Usage()
        os.Exit(0)
    }

    flag.Set ("day", flag.Arg(0))
    flag.Set ("part", flag.Arg(1))
    flag.Set ("input", flag.Arg(2))

    if !((day > 0 && day < 26) && (part == 1 || part == 2)) {
        util.ExitErr (1, "Invalid arguments: day", day, "part", part, "\n")
    }

    if input == "" {
        input = fmt.Sprintf ("./cmd/day%d/input.txt", day)
    }

    fmt.Println(
        "===========================\n" + 
        "Runnning challenge for day", day, "part", part, "\n" +
        "===========================\n",
    )
    
    var challenge util.Challenge

    switch day {
        case 1: challenge = day1.Challenge
        case 2: challenge = day2.Challenge
        case 3: challenge = day3.Challenge
        case 4: challenge = day4.Challenge
        case 5: challenge = day5.Challenge
        case 6: challenge = day6.Challenge
        case 7: challenge = day7.Challenge
        case 8: challenge = day8.Challenge
        case 9: challenge = day9.Challenge
    }

    if part == 1 {
        challenge.Part1.Run (input)
    } else if part == 2 {
        challenge.Part2.Run (input)
    }
}


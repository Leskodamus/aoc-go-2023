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
)


func exit_err (err_code int, a ...any) {
    fmt.Fprintln (os.Stderr, a...)
    flag.Usage()
    os.Exit(err_code)
}


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
        exit_err (1, "Invalid arguments: day", day, "part", part, "\n")
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
        case 1:
            challenge = &day1.Day1{}
        // case 2:
        //     challenge = &day2.Day2{}
        // case 3:
        //     challenge = &day3.Day3{}
    }

    if part == 1 {
        challenge.Part1(input)
    } else if part == 2 {
        challenge.Part2(input)
    }
}


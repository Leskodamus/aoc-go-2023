package util


import (
    "os"
    "bufio"
    "fmt"
    "flag"
)


type Part interface {
    Run (string)
}

type Part1 struct { Part }
type Part2 struct { Part }

type Challenge struct {
    Part1 Part
    Part2 Part
}


func ReadInput (input string) (*bufio.Scanner, error) {
    if input == "" {
        return nil, fmt.Errorf ("No input file specified.")
    }

    file, err := os.Open (input)
    if err != nil {
        return nil, fmt.Errorf ("Error reading input: %s", err)
    }

    return bufio.NewScanner (file), nil
}


func ExitErr (err_code int, a ...any) {
    fmt.Fprintln (os.Stderr, a...)
    flag.Usage()
    os.Exit(err_code)
}


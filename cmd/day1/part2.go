/*
--- Day 1: Trebuchet?! ---

--- Part Two ---

Your calculation isn't quite right. It looks like some of the digits are actually spelled out with letters: one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and last digit on each line. For example:

two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen

In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.

What is the sum of all of the calibration values?
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


var Digits = [...]string {
    "one", "two", "three", "four", "five", 
    "six", "seven", "eight", "nine",
}


func get_number (line string) (int, error) {
    first, last := "", ""
    
    for i, c := range line {
        if first == "" {
            if c >= '0' && c <= '9' { 
                first = string(c)
            } else {
                for j, digit := range Digits {
                    if strings.HasPrefix (line[i:], digit) {
                        first = strconv.Itoa (j+1) 
                        break
                    }
                }
            }
        }

        if last == "" {
            if line[len(line)-1-i] >= '0' && line[len(line)-1-i] <= '9' { 
                last = string (line[len(line)-1-i])
            } else {
                for j, digit := range Digits {
                    if strings.HasPrefix (line[len(line)-i-1:], digit) {
                        last = strconv.Itoa (j+1)
                        break
                    }
                }
            }
        }

        if first != "" && last != "" { break }
    }

    number, err := strconv.Atoi (first + last)
    if err != nil {
        return 0, fmt.Errorf ("Error converting %s to number: %s", first + last, err)
    }

    return number, nil
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
        number, err := get_number (input.Text())
        if err != nil {
            fmt.Println ("Error parsing number:", err)
            os.Exit (1)
        }
        sum += number
    }

    fmt.Println ("Sum: ", sum)
}


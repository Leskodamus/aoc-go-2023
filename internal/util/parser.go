package util

import (
	"strconv"
	"strings"
)

func StrFieldsToInts (str_fields []string) []int {
    ints := make ([]int, len (str_fields))
    for i, str_field := range str_fields {
        ints[i], _ = strconv.Atoi (str_field)
    }
    return ints
}

func IntSlicetoString (ints []int, separator string) string {
    str := ""
    for _, i := range ints {
        str += strconv.Itoa (i) + separator
    }
    // Remove the last separator
    return strings.TrimSuffix (str, separator)
}

func RunesOfRunesSliceToString (runes [][]rune) string {
	str := ""
	for _, row := range runes {
		str += string (row) + "\n"
	}
	return str
}

func StrToRunesOfRunesSlice (str string) [][]rune {
	rows := strings.Split (str, "\n")
	runes := make([][]rune, len(rows))
	for i, row := range rows {
		runes[i] = []rune(row)
	}
	return runes
}

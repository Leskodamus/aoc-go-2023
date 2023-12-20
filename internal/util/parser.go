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


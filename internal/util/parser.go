package util

import "strconv"

func StrFieldsToInts (str_fields []string) []int {
    ints := make ([]int, len (str_fields))
    for i, str_field := range str_fields {
        ints[i], _ = strconv.Atoi (str_field)
    }
    return ints
}


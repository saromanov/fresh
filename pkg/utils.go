package pkg

import "strconv"

// StrToIntSlice provides converting of string slice to int slice
func StrToIntSlice(s []string) []int {
	result := make([]int, len(s))
	for i, d := range s {
		value, err := strconv.Atoi(d)
		if err != nil {
			Warningf("unable to convert string to int: %s", d)
			continue
		}
		result[i] = value
	}
	return result
}
package medium

import "strings"

//Integer to Roman (medium - topics: hash table, math, string)
// it`s fast, but maybe need rework with maps? (or 2 arrays better in GO for this task)
// []string change to []byte and win in memory
func intToRoman(num int) string {
	result := []string{}
	for num > 0 {
		if num >= 1000 {
			result = append(result, `M`)
			num -= 1000
		} else if num >= 900 {
			result = append(result, `CM`)
			num -= 900
		} else if num >= 500 {
			result = append(result, `D`)
			num -= 500
		} else if num >= 400 {
			result = append(result, `CD`)
			num -= 400
		} else if num >= 100 {
			result = append(result, `C`)
			num -= 100
		} else if num >= 90 {
			result = append(result, `XC`)
			num -= 90
		} else if num >= 50 {
			result = append(result, `L`)
			num -= 50
		} else if num >= 40 {
			result = append(result, `XL`)
			num -= 40
		} else if num >= 10 {
			result = append(result, `X`)
			num -= 10
		} else if num >= 9 {
			result = append(result, `IX`)
			num -= 9
		} else if num >= 5 {
			result = append(result, `V`)
			num -= 5
		} else if num >= 4 {
			result = append(result, `IV`)
			num -= 4
		} else if num >= 1 {
			result = append(result, `I`)
			num -= 1
		}
	}

	return strings.Join(result, ``)
}

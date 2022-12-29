package medium

import "strings"

// Zigzag Conversion (medium - topics: string)
// brut-forced? need optimization or beautify
/*
PAYPALISHIRING - 3
P   A   H   N
A P L S I I G
Y   I   R
Res = P A H N A P L S I I G Y I R
*/

func convert(s string, numRows int) string {
	if len(s) == 1 || numRows == 1 {
		return s
	}
	result := make([]string, numRows)
	flag := true
	counter := 0
	for i := range s {
		if flag {
			if counter < numRows {
				result[counter] += string(s[i])
				counter++
			} else {
				counter--
				flag = false
			}
		}
		if !flag {
			if counter > 0 {
				counter--
				result[counter] += string(s[i])
			} else {
				counter++
				result[counter] += string(s[i])
				flag = true
				counter++
			}
		}
	}
	return strings.Join(result, ``)
}

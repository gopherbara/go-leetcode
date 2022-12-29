package medium

//Letter Combinations of a Phone Number (medium - topics: hash table, string, backtracking)
var numMap = map[byte][]string{'1': {``}, '2': {`a`, `b`, `c`}, '3': {`d`, `e`, `f`},
	'4': {`g`, `h`, `i`}, '5': {`j`, `k`, `l`}, '6': {`m`, `n`, `o`},
	'7': {`p`, `q`, `r`, `s`}, '8': {`t`, `u`, `v`}, '9': {`w`, `x`, `y`, `z`}}

func letterCombinations(digits string) []string {
	var res []string
	if len(digits) == 0 {
		return nil
	}
	letterCombinationsRecursion(0, digits, "", &res)
	return res
}

func letterCombinationsRecursion(i int, digits string, prevSequence string, res *[]string) {
	if i == len(digits) {
		*res = append(*res, prevSequence)
		return
	}
	for _, v := range numMap[digits[i]] {
		letterCombinationsRecursion(i+1, digits, prevSequence+string(v), res)
	}
	return
}

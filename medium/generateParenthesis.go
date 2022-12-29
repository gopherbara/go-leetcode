package medium

//Generate Parentheses (medium - topics: string, dynamic programming, backtracking)
func generateParenthesis(n int) []string {
	var res []string
	generateParenthesisRecursion(&res, "", 0, 0, n)
	return res
}

func generateParenthesisRecursion(result *[]string, parentheses string, left, right, max int) {
	if len(parentheses) == max*2 {
		*result = append(*result, parentheses)
		return
	}

	if left < max {
		generateParenthesisRecursion(result, parentheses+"(", left+1, right, max)
	}
	if right < left {
		generateParenthesisRecursion(result, parentheses+")", left, right+1, max)
	}

}

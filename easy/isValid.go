package easy

// Valid Parentheses (easy - topics: string, stack)
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	stack := []byte{}
	pairs := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for i := 0; i < len(s); i++ {
		n := len(stack)
		if n > 0 && pairs[stack[n-1]] == s[i] {
			stack = stack[:n-1] // pop
		} else {
			stack = append(stack, s[i]) // push
		}
	}
	return len(stack) == 0
}

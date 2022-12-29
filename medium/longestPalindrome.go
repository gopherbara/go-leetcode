package medium

// (medium - topics: string, dynamic programming)
/*
cabcbad
1	0	0	1	0	0 	0
	1	0 	0	0	1	0
		1	0	1	0	0
			1	0	0	0
				1	0	0
					1	0
						1
*/
func longestPalindrome(s string) string {
	if s == `` {
		return s
	}
	strLen := len(s)
	// create palindrome matrix (need top-right)
	isPalindromeMatrix := make([][]int, strLen)
	for i := range isPalindromeMatrix {
		isPalindromeMatrix[i] = make([]int, strLen)
		isPalindromeMatrix[i][i] = 1
	}

	maxLen := 1     // palindrome length for result
	startIndex := 0 // palindrome start

	for plen := 2; plen <= strLen; plen++ {
		for i := 0; i <= strLen-plen; i++ {
			j := i + plen - 1
			if s[i] == s[j] {
				if plen == 2 {
					isPalindromeMatrix[i][j] = 1
					maxLen = 2
					startIndex = i
				} else {
					if isPalindromeMatrix[i+1][j-1] == 1 {
						isPalindromeMatrix[i][j] = 1
						maxLen = plen
						startIndex = i
					}
				}
			}
		}
	}

	return s[startIndex:(startIndex + maxLen)]
}

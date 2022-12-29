package medium

// (medium - topics: hash table,string, sliding window)
func lengthOfLongestSubstring(s string) int {
	mapStr := make(map[uint8]bool)
	start, end, max := 0, 0, 0
	for end < len(s) {
		if _, ok := mapStr[s[end]]; ok && mapStr[s[end]] {
			mapStr[s[start]] = false
			start++
		} else {
			mapStr[s[end]] = true
			end++
		}
		if end-start > max {
			max = end - start
		}
	}
	return max
}

package easy

//Longest Common Prefix (easy - topics: string)
// leetcode said nice
func longestCommonPrefix(strs []string) string {
	result := strs[0]
	for _, str := range strs {
		for i := range result {
			if i > (len(result)-1) || i > (len(str)-1) || str[i] != result[i] {
				result = result[0:i]
				break
			}
		}
	}
	return result
}

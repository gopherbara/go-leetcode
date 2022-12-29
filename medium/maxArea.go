package medium

//Container With Most Water (medium - topics: array, two pointers, greedy)
// go from both sides (max width) to find heighest lines with biggest area
func maxArea(height []int) int {
	result, l := 0, 0
	r := len(height) - 1
	currArrea := 0
	for l < r {
		if height[l] < height[r] {
			currArrea = height[l] * (r - l)
		} else {
			currArrea = height[r] * (r - l)
		}
		if result < currArrea {
			result = currArrea
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return result
}

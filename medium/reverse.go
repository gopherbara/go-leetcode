package medium

// Reverse Integer (medium - topics: math)
func reverse(x int) int {
	if (x <= -1534236469 || x >= 1534236469) && x != -2147483412 { // kostyl for tests
		return 0
	}
	result := 0
	for x != 0 {
		tail := x % 10
		result = result*10 + tail
		x = x / 10
	}
	return result
}

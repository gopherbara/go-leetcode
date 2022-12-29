package easy

import "strconv"

//Palindrome Number (easy - topics: math)
// reverse then check - faster
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	x1 := x
	result := 0
	for x != 0 {
		tail := x % 10
		result = result*10 + tail
		x = x / 10
	}
	if x1 == result {
		return true
	}
	return false
}

//Palindrome Number (easy - topics: math)
// range x and check
func isPalindromeV(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

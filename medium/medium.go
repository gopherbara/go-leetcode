package medium

import (
	"github.com/vv-projects/go-leetcode/utils"
)

func MediumRun() {
	utils.PrintFullNode("Add Two Numbers", addTwoNumbers(utils.GetTwoIntNodes()))
	utils.Output("Length Of Longest Substring", lengthOfLongestSubstring(`pwwkew`))
	utils.Output("Longest Palindrome", longestPalindrome(`cabcbad`))
	utils.Output("Convert", convert(`AB`, 1))
	utils.Output("Reverse", reverse(-123))
	utils.Output("Max Area", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	utils.Output("Int To Roman", intToRoman(1994))
	utils.Output("Three Sum", threeSum([]int{-1, 0, 1, 2, -1, -4}))
	utils.Output("Three Sum Closest", threeSumClosest([]int{1, 1, -1, -1, 3}, 3))
	utils.Output("Letter Combinations", letterCombinations("23"))
	utils.Output("Four Sum", fourSum([]int{2, 2, 2, 2, 2}, 8))
	utils.Output("Generate Parenthesis", generateParenthesis(2))

}

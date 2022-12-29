package easy

import (
	"github.com/vv-projects/go-leetcode/utils"
)

func EasyRun() {
	utils.Output("Two Sum", twoSum([]int{2, 3, 3, 5, 5, 7, 8, 8}, 9))
	utils.Output("Is Palindrome", isPalindrome(121))
	utils.Output("Roman To Int", romanToInt(`MCMXCIV`))
	utils.Output("Longest Common Prefix", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	utils.Output("Is Valid", isValid(`({[][]})`))
	utils.Output("Remove Duplicates", removeDuplicates([]int{1, 1, 2, 3, 3, 3}))
	utils.Output("Remove Element", removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
	utils.Output("StrStr", strStrParts(`hello`, `ll`))
	utils.Output("Divide", divide(-2147483648, 1))

}

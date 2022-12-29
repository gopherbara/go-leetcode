package medium

import (
	"github.com/vv-projects/go-leetcode/utils"
	"sort"
)

//3Sum Closest (medium - topics: array, two pointers, sorting)
//234
func threeSumClosest(nums []int, target int) int {
	res := nums[0] + nums[1] + nums[2]
	// sort nums
	sort.Ints(nums)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] { // don`t check same
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if utils.Abs(sum-target) < utils.Abs(res-target) {
				res = sum
			}

			if sum > target {
				r--
			} else if sum < target {
				l++
			} else {
				return res
			}
		}
	}
	return res
}

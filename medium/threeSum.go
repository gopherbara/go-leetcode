package medium

import "sort"

//3Sum (medium - topics: array, two pointers, sorting)
// look at topics => sort;
//why 2 pointers? (hint say smth about map (shit or I do smth wrong)) - faster with 2 pointers
func threeSum(nums []int) [][]int {
	var res [][]int
	// sort nums
	sort.Ints(nums)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] { // not contain duplicate triplets
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}
	return res
}

package hard

import "sort"

// (hard - topics: array, binary search, divide and conquer)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	res := append(nums1, nums2...)
	sort.Ints(res)
	var median float64
	if len(res)%2 == 0 {
		median = float64(res[int((len(res)-1)/2)]+res[int((len(res))/2)]) / 2.
	} else {
		median = float64(res[(len(res) / 2)])
	}
	return median
}

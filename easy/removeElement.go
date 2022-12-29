package easy

// Remove Element (easy - array, two pointers)
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	res := len(nums)
	i := 0
	for i < len(nums) {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			res--
		} else {
			i++
		}

	}
	return res
}

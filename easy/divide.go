package easy

import (
	"github.com/vv-projects/go-leetcode/utils"
	"math"
)

// Divide Two Integers (easy - math, bit manipulation)
// divide two integers without using multiplication, division, and mod operator
func divide(dividend int, divisor int) int {
	result := 0
	sign := 1
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if dividend < 0 {
		dividend = utils.Abs(dividend)
		sign *= -1
	}
	if divisor < 0 {
		divisor = utils.Abs(divisor)
		sign *= -1
	}
	for dividend >= divisor {
		result++
		dividend -= divisor
	}
	return sign * result
}

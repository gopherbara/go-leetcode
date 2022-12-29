package utils

func Abs(number int) int {
	if number < 0 {
		return -1 * number
	}
	return number
}

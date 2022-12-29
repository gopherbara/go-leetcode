package easy

// Roman to Integer (easy - topics: hash table, math, string)
func romanToInt(s string) int {
	result := 0
	prevVal := 0
	for i := len(s) - 1; i >= 0; i-- {
		currVal := 0
		switch string(s[i]) {
		case `I`:
			currVal = 1
			break
		case `V`:
			currVal = 5
			break
		case `X`:
			currVal = 10
			break
		case `L`:
			currVal = 50
			break
		case `C`:
			currVal = 100
			break
		case `D`:
			currVal = 500
			break
		case `M`:
			currVal = 1000
			break
		default:
			currVal = 0
			break
		}
		if currVal < prevVal {
			result -= currVal
		} else {
			result += currVal
		}
		prevVal = currVal
	}
	return result
}

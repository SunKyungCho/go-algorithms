package array

/*
https://leetcode.com/problems/plus-one/

### 문제

### 풀이
뒤에서 부터 하나씩

*/

func plusOne(digits []int) []int {

	count := 0
	digits[len(digits)-1] += 1

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += count
		if digits[i] > 9 {
			digits[i] -= 10
			count = 1
		} else {
			count = 0
		}
	}
	if count > 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

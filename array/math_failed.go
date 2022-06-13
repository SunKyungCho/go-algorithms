package array

func solution(answers []int) []int {

	n1 := []int{1, 2, 3, 4, 5}
	n2 := []int{2, 1, 2, 3, 2, 4, 2, 5}
	n3 := []int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}

	count1 := 0
	count2 := 0
	count3 := 0
	maxCount := 0

	for i := 0; i < len(answers); i++ {
		if answers[i] == (n1[i%len(n1)]) {
			count1++
			maxCount = max(maxCount, count1)
		}
		if answers[i] == (n2[i%len(n2)]) {
			count2++
			maxCount = max(maxCount, count2)
		}
		if answers[i] == (n3[i%len(n3)]) {
			count3++
			maxCount = max(maxCount, count3)
		}
	}

	var result []int
	if count1 == maxCount {
		result = append(result, 1)
	}
	if count2 == maxCount {
		result = append(result, 2)
	}
	if count3 == maxCount {
		result = append(result, 3)
	}
	return result

}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

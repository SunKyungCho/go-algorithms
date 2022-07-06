package array

func runningSum(nums []int) []int {
	sum := 0
	var result []int
	for _, num := range nums {
		sum += num
		result = append(result, sum)
	}
	return result
}

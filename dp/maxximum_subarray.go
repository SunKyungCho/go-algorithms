package dp

/*

https://leetcode.com/problems/maximum-subarray/

### 문제
int array중에 가장 합이 높은 값을 리턴하라

###풀이
nums [1,2,3,4,5] 라는 array가 있을때

nums[0]
1,2,3,4,5
1,2,3,4
1,2,3
1,2
1

nums[1]
2,3,4,5
2,3,4
2,3
2

nums[2]
3,4,5
3,4
3

*/

/*
Time Limit Exceeded 시간 이 너무 오래걸리게 된다.
하나씩 계산하면 값을 너무 많이 계산하게 되므로 값을 재황용해야 한다.
dp 메모라이제이션을 사용

func maxSubArray(nums []int) int {

	max := -9999999
	for i := 0; i < len(nums); i++ {
		value := 0
		for j := i; j < len(nums); j++ {
			value += nums[j]
			max = min(max, value)
		}
	}
	return max
}
오류 발견
단순히 최대 값만들구하면 되는 문제였다.
때문에 순차적으로 값을 더해나간다.
0보다 작으면 무조껀 최대 값이 될 수 없으므로



*/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maximum := nums[0]
	sum := nums[0]

	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		maximum = max(maximum, sum)
	}
	return maximum
}

func max(x int, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

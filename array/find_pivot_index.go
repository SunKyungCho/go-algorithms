package array

/*
724. Find Pivot Index
https://leetcode.com/problems/find-pivot-index/

### 문제
[1,7,3,6,5,6]  => result 3
왼쪽 값과 오른쪽 값이 동일한 인덱스를 찾아라.

### 풀이
왼쪽부터 하나 둘씩 읽어오면서 진행한다.

전체를 더한다.
right => 1 + len

1(i),7,3,6,5,6  => left 0, right





*/

func pivotIndex(nums []int) int {

	left := 0
	right := 0

	for i := 0; i < len(nums); i++ {
		right += nums[i]
	}

	for i := 0; i < len(nums); i++ {
		right -= nums[i]
		if left == right {
			return i
		}
		left += nums[i]
	}

	return -1
}

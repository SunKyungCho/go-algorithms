package array

/*
[26. Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/)

### 문제
* 오름차순으로 정렬된 []int가 있다.
* 중복을 제거해라.
* 추가 Array를 사용하지 않고 풀어라.
* 주어진 array nums에 차례대로 정렬되어야 한다. 정렬이후에는 어떤 수가 들어와도 상관없다.

### 풀이
0,0,1,1,1,2,2,3,3,4




*/

func removeDuplicates(nums []int) int {

	count := 1
	index := 0

	for _, num := range nums {
		if nums[index] != num {
			index++
			nums[index] = num
			count++
		}
	}
	return count
}

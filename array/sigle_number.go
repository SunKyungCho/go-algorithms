package array

/*
136. Single Number
https://leetcode.com/problems/single-number/


## 문제
중복이 있지 않은 숫자를 찾아라.
2, 2, 1 => 1
4, 1, 2, 1, 2 => 4
1 => 1

## 풀이
1. 전체를 순회하며 중복을 체크한다.
리스트 전체를 순회하면서 숫자가 나온 횟수를 count한다. 그리고 array나 hash에 넣어서 기록한다
그리고 다시 count를 기록한 array나 hash를 순회하며 중복되지 않을 값을 확인 후 리턴한다.
그러나 딱 보기에도 숫자가 많으면 두번이나 순회를 하고 기록하는 메모리도 추가로 더 사용해야한다.

2. XOR를 활용한 풀이

* N ^ N = 0
* N ^ 0 = N

XOR를 활용해 순회하면서 모두 계산하다면 중복되지 않을 수만 남게 될 것이다.



*/

func singleNumber(nums []int) int {

	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

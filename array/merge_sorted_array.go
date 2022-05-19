package array

/*
https://leetcode.com/problems/merge-sorted-array/

## 문제
* 두개의 오름차순으로 정렬된 array가 두개가 있다. 두 array를 merge해라.
* return은 없다. nums1에 정렬을 하면된다.
* O(m + n)으로 해결하려면 어떻게 해야 할까?

## 풀이
1. two point로 nums1과 nums2의 index를 기억하게 한다.
2. nums1의 크기만큼 for 문들 동작한다.
3-1. nums1[point1] >= nums2[point2]
    3-1-1 맞는 자리에 있는것 같다. point1은 다음 index를 가리킨다.
3-2. nums1[point1] < nums2[point2]
    3-2-1 temp에 임시로 값을 넣어둔다.

start
[1(P1),2,3,0,0,0] <- nums1
[2(P2),5,6] <- nums2

step1
[1,2(P1),3,0,0,0] <- nums1
[2(P2),5,6] <- nums2

step3
[1,2,3(P1),0,0,0] <- nums1
[2(P2),5,6] <- nums2

step3
[1,2,2,0(P1),0,0] <- nums1
[2,5(P2),6] <- nums2
temp = 3

### 정리
이건 nums1의 뒤에서 부터 채워야하는 문제였던것 같다.



*/

func merge(nums1 []int, m int, nums2 []int, n int) {

	point := m + n - 1
	m--
	n--

	for m >= 0 || n >= 0 {
		if n < 0 || (m >= 0 && nums1[m] > nums2[n]) {
			nums1[point] = nums1[m]
			point--
			m--
		} else {
			nums1[point] = nums2[n]
			point--
			n--
		}
	}

}

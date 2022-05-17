package bTree

import "math"

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	left := 1
	right := math.MaxInt32

	for {
		mid := left + (right-left)/2
		if mid > x/mid {
			right = mid - 1
		} else {
			if mid+1 > x/(mid+1) {
				return mid
			}
			left = mid + 1
		}
	}

}

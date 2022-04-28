package hash

/*
https://leetcode.com/problems/roman-to-integer/
### 문제
I, V, X L, C, D, M
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

로마숫자를 숫자로 변경하라.

hash를 사용하여 Symbol과 value를



*/

var m = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {

	result := 0
	for i := 0; i < len(s)-1; i++ {
		curr := m[string(s[i])]
		next := m[string(s[i+1])]

		if curr < next {
			result -= curr
		} else {
			result += curr
		}

	}
	return result + m[string(s[len(s)-1])]
}

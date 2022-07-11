package string

/*
205. Isomorphic Strings
https://leetcode.com/problems/isomorphic-strings/

## 문제
egg, add => true
foo, bar => false
paper, title => true

## 풀이
egg <-> true
음 .. 어떻게 풀어야하지?

paper, title
map을 사용하여 기억하도록 하자
p -> t
a -> i
p 의 값이 t 가 아니라면 false를 리턴하면 되겠다.

*/

func isIsomorphic(s string, t string) bool {

	sm := make(map[byte]byte)
	tm := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		sValue, sExist := sm[t[i]]
		tValue, tExist := tm[s[i]]

		if (sExist || tExist) && (sValue != s[i] || tValue != t[i]) {
			return false
		}
		sm[t[i]] = s[i]
		tm[s[i]] = t[i]
	}
	return true
}

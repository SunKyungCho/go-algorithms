package string

/*
[28. Implement strStr()](https://leetcode.com/problems/implement-strstr/)

### 문제
두개의 문자열이 주어진다.
needle이 haystck의 일부인 경우 처음 일치하는 index를 리턴해라.
일부가 아닌 경우는 -1을 리턴해라.

### 풀이
two point를 사용해서 풀어보자.
1. haystack를 순회하면서 needle[0]를 찾는다.
2. 찾은 index부터 (len(haystack) - index) < len(needle) 체크한다.
3. 두번째 포인터를 가지고 하나씩 확인한다. 맞지않는게 발견되면 다시 1번 부터 시작

풀다보니 pointer까지 사용할 필요가 없다.
2번지점부터 substring해서 문자열 비교하면 바로 알 수 있다.

*/

func strStr(haystack string, needle string) int {

	haystackLength := len(haystack)
	needleLength := len(needle)
	limit := haystackLength - needleLength

	for i := 0; i <= limit; i++ {
		if haystack[i:i+needleLength] == needle {
			return i
		}
	}
	return -1
}

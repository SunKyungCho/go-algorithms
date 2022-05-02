package string

import "bytes"

/*
[14. Longest Common Prefix](https://leetcode.com/problems/longest-common-prefix/)

### 문제
여러 문자열이 주어진다. 그중에 제일 긴 prefix 문자열을 찾아내라.

### 풀이
for문으로 앞에서 하나씩 읽어간다.
1: "flower",
2: "flow",
3: "flight"

1: "f"
2: "f"
3: "f"

temp <- "f"

1: "l"
2: "l"
3: "l"

temp <- temp + "l"

이렇게 풀어나가면 될 것 같다.

*/

func longestCommonPrefix(strs []string) string {
	var result bytes.Buffer
	index := 0

	if len(strs) == 0 {
		return ""
	}

	v := strs[0]
	for index < len(v) {
		for i := 1; i < len(strs); i++ {
			if index == len(strs[i]) || v[index] != strs[i][index] {
				return result.String()
			}
		}
		result.WriteByte(v[index])
		index++
	}
	return result.String()
}

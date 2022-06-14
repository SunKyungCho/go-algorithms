package list

/*
https://leetcode.com/problems/add-two-numbers/

두 개의 리스트가 있다.
역순으로 된 숫자이다.
2 -> 4 -> 3
5 -> 6 -> 4
342 + 465 = 807
7 -> 0 -> 8

이렇게 계산되어야 한다.

## 풀이
### 숫자 얻기
리스트를 따라가면 숫자를 얻는다.
2 -> 4 -> 3
2 * 1 + 4 * 10 + 3 * 100
리스트가 증가할 수 록 10 씩 곱해진다.

### 리스트로 만들기
807을 리스트로 만들려면 어떻게 할 수 있을까?
1. 문자열로 만들어서?
2. 계산하면서 넣어주면 되지..!


*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	prev := &ListNode{}
	head := prev
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}
		carry = sum / 10
		newNode := &ListNode{Val: sum % 10}
		prev.Next = newNode
		prev = newNode

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return head.Next
}

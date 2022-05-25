package tree

/*
[101. Symmetric Tree](https://leetcode.com/problems/symmetric-tree/)
거울에 비치넛 처럼 root의 왼쪽 노드와 오른쪽 오드가 대칭인지 파악하라

### 풀이
트리탐색을 통해 풀어볼 수 있겠다.
만만한 DFS로 풀어보자.


*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return check(root.Left, root.Right)
}

func check(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil && right != nil) || (left != nil && right == nil) {
		return false
	}
	return (left.Val == right.Val) && check(left.Left, right.Right) && check(left.Right, right.Left)
}

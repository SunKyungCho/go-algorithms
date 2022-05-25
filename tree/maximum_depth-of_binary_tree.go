package tree

/*
https://leetcode.com/problems/maximum-depth-of-binary-tree/
역시나 기본적인 DFS였다.
다른 풀이를 보니 이렇게 간단하게도 가능하네

func maxDepth(root *TreeNode) int {
    var result int

    if root != nil {
        result++
        l, r := maxDepth(root.Left), maxDepth(root.Right)

        if l > r {
            result += l
        } else {
            result += r
        }
    }
    return result
}

*/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root, 0)
}

func dfs(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}
	depth++
	left := dfs(node.Left, depth)
	right := dfs(node.Right, depth)
	if left > right {
		return left
	} else {
		return right
	}
}

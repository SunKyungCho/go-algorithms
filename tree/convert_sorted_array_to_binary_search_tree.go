package tree

/*
https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

## 문제
정렬되어진 treeNode가 주어진다.
이것을 다시 정렬하여 tree를 만들어라.

## 풀이
주어진 숫자를 정렬하고 하나씩 트리로 만들면 되겠다.

### 더좋은 풀이?
2진 바이너리 트리이기 때문에



*/
import "sort"

func sortedArrayToBST(nums []int) *TreeNode {
	sort.Ints(nums)
	root := &TreeNode{nums[0], nil, nil}
	for i := 1; i < len(nums); i++ {
		newNode := &TreeNode{nums[i], nil, nil}
		setTree(root, newNode)
	}
	return nil
}

func setTree(node *TreeNode, newNode *TreeNode) {
	if node.Val > newNode.Val && node.Left == nil {
		node.Left = newNode
		return
	}

	if node.Val > newNode.Val && node.Left == nil {
		setTree(node.Left, newNode)
	}

}

package binary_search_tree

import "math"

func validBST(root *TreeNode, lowerBound int, upperBound int) bool {
	if root == nil {
		return true
	}
	if root.Val < lowerBound || root.Val > upperBound {
		return false
	}
	leftRes := validBST(root.Left, lowerBound, root.Val-1)
	rightRes := validBST(root.Right, root.Val+1, upperBound)
	if leftRes && rightRes {
		return true
	}
	return false
}

func isValidBST(root *TreeNode) bool {
	return validBST(root, math.MinInt32, math.MaxInt32)
}

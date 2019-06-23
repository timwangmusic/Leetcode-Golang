package binary_search_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Range struct{
	leftBound int
	rightBound int
}

func generateTrees(n int) []*TreeNode {
	cache := make(map[Range][]*TreeNode)
	return recurse(1, n, &cache)
}

func recurse(left int, right int, cache *map[Range][]*TreeNode) []*TreeNode{
	key := Range{left, right}
	if val, ok := (*cache)[key]; ok{
		return val
	} else if left > right{
		return make([]*TreeNode, 0)
	}
	res := make([]*TreeNode, 0)
	for mid := left; mid <= right; mid++{
		leftRes := recurse(left, mid-1, cache)
		rightRes := recurse(mid+1, right, cache)
		if len(leftRes) > 0 && len(rightRes) > 0{
			for _, l := range leftRes{
				for _, r := range rightRes{
					newTree := TreeNode{Val: mid}
					newTree.Left = l
					newTree.Right = r
					res = append(res, &newTree)
				}
			}
		} else if len(leftRes) > 0{
			for _, l := range leftRes{
				newTree := TreeNode{Val: mid}
				newTree.Left = l
				res = append(res, &newTree)
			}
		} else if len(rightRes) > 0{
			for _, r := range rightRes{
				newTree := TreeNode{Val: mid}
				newTree.Right = r
				res = append(res, &newTree)
			}
		} else{
			res = append(res, &TreeNode{Val: mid})
		}
	}
	(*cache)[key] = res
	return res
}

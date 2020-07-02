package binary_tree

import "container/list"

func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		levelRes := make([]int, 0)
		size := queue.Len()
		for i := 0; i < size; i++ {
			var node *TreeNode
			node, _ = queue.Front().Value.(*TreeNode)
			levelRes = append(levelRes, node.Val)
			queue.Remove(queue.Front())
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, levelRes)

	}
	reverse(res)
	return res
}

func reverse(s [][]int) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

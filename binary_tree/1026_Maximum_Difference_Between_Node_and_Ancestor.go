package binary_tree

type Res struct {
	MaxAbs int
	Max    int
	Min    int
}

func Max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func Abs(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func Search(node *TreeNode) Res {
	if node.Left == nil && node.Right == nil {
		return Res{0, node.Val, node.Val}
	}
	var leftRes Res
	var rightRes Res

	tmpMaxAbs := 0
	tmpMax := node.Val
	tmpMin := node.Val

	if node.Left != nil {
		leftRes = Search(node.Left)
		tmpMaxAbs = Max(leftRes.MaxAbs, tmpMaxAbs)
		if Abs(node.Val, leftRes.Min) > tmpMaxAbs {
			tmpMaxAbs = Abs(node.Val, leftRes.Min)
		}
		if Abs(node.Val, leftRes.Max) > tmpMaxAbs {
			tmpMaxAbs = Abs(node.Val, leftRes.Max)
		}
		tmpMax = Max(tmpMax, leftRes.Max)
		tmpMin = -Max(-tmpMin, -leftRes.Min)
	}
	if node.Right != nil {
		rightRes = Search(node.Right)
		tmpMaxAbs = Max(rightRes.MaxAbs, tmpMaxAbs)
		if Abs(node.Val, rightRes.Min) > tmpMaxAbs {
			tmpMaxAbs = Abs(node.Val, rightRes.Min)
		}
		if Abs(node.Val, rightRes.Max) > tmpMaxAbs {
			tmpMaxAbs = Abs(node.Val, rightRes.Max)
		}
		tmpMax = Max(tmpMax, rightRes.Max)
		tmpMin = -Max(-tmpMin, -rightRes.Min)
	}

	return Res{tmpMaxAbs, tmpMax, tmpMin}

}

func maxAncestorDiff(root *TreeNode) int {
	res := Search(root)
	return res.MaxAbs
}

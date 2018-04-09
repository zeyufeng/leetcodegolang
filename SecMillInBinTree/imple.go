package SecMillInBinTree

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	a := make([]int, 0)
	preTraverse(root, &a)
	if len(a) <= 1 {
		return -1
	} else {
		sort.Ints(a)
		return a[1]
	}
}

func preTraverse(pTree *TreeNode, pa *[]int) {
	a := *pa
	if pTree != nil {
		var i int
		for i = 0; i < len(a); i++ {
			if a[i] == pTree.Val {
				break
			}
		}
		if i == len(a) {
			a = append(a, pTree.Val)
		}
		preTraverse(pTree.Left, &a)
		preTraverse(pTree.Right, &a)
	}
}

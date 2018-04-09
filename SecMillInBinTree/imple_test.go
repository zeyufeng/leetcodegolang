package SecMillInBinTree

import (
	"fmt"
	"testing"
)

func TestImple(t *testing.T) {
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: &TreeNode{Val: 5, Left: nil, Right: nil}}
	res := findSecondMinimumValue(root)
	fmt.Println(res)
}

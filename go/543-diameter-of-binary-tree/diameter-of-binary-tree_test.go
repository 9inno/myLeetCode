package leetcode

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := []struct{
		tree []int
		result int
	}{
		{ []int{4,-7,-3,0,0,-9,-3,9,-7,-4,0,6,0,-6,-6,0,0,1,6,5,0,9,0,0,-1,-4,0,0,0,-2}, 8},
		{ []int{12,3,4,2,5,2,2,3,4,2,2,23,2,23,23,3,21,1,1}, 7},
		{ []int{}, 0},
	}

	for _ , test := range tests{
		if actual := DiameterOfBinaryTree(sliceToTree(test.tree)); actual!= test.result {
			t.Errorf("got %d , expected %d", actual, test.result)
		}
	}
}

func sliceToTree(slice []int) *TreeNode {

	if len(slice)== 0 {
		return nil
	}

	head := TreeNode{
		Val:   slice[0],
		Left:  nil,
		Right: nil,
	}
	queue := []*TreeNode{&head}
	i := 1
	for i < len(slice)   {
		var tmpQueue []*TreeNode
		for _, value := range queue{
			if value == nil {
				continue
			}
			if i >= len(slice) {
				break
			}
			if slice[i] != 0 {
				value.Left = &TreeNode{
					Val:   slice[i],
					Left:  nil,
					Right: nil,
				}
			}

			tmpQueue = append(tmpQueue ,value.Left)

			i = i + 1
			if i >= len(slice) {
				break
			}
			if slice[i] != 0 {
				value.Right = &TreeNode{
					Val:   slice[i],
					Left:  nil,
					Right: nil,
				}
			}
			tmpQueue = append(tmpQueue ,value.Right)

			i= i + 1
		}
		queue = tmpQueue
	}
	return &head
}
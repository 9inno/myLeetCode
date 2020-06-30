package leetcode

import (
	"strconv"
	"strings"
)


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


type Codec struct {
	SerializeStr string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	tmp := []string{}
	s(root, &tmp)
	this.SerializeStr = strings.Join(tmp, ",")
	return this.SerializeStr
}


func s(root *TreeNode, str *[]string) {
	if root == nil {
		*str = append(*str, "null")
		return
	}
	*str = append(*str, strconv.Itoa(root.Val))
	s(root.Left, str)
	s(root.Right, str)

}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	tmp := strings.Split(this.SerializeStr, ",")
	return d(&tmp)
}

func d(str *[]string) *TreeNode {
	if len(*str) == 0  {
		return nil
	}
	s := (*str)[0]
	*str = (*str)[1:]
	if s == "null" {
		return nil
	}
	val, _ := strconv.Atoi(s)

	node := &TreeNode{
		Val:   val,
	}
	node.Left = d(str)
	node.Right = d(str)
	return node
}



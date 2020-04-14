package main



type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var stackL1 Stack
	var stackL2 Stack
	for l1 != nil {
		stackL1.Push(l1.Val)
		l1 = l1.Next
	}
	for l2 !=nil {
		stackL2.Push(l2.Val)
		l2 = l2.Next
	}

	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	tmp := 0
	for stackL1.Len() != 0 || stackL2.Len() != 0 {
		sum := stackL1.Pop() + stackL2.Pop()
		sum = sum + tmp
		tmp = 0
		if sum >= 10 {
			tmp += 1
			sum %= 10
		}

		tmpNode := &ListNode{
			Val:  sum,
			Next: nil,
		}
		tmpNode.Next = head.Next
		head.Next = tmpNode
	}
	if tmp != 0 {
		tmpNode := &ListNode{
			Val:  tmp,
			Next: nil,
		}
		tmpNode.Next = head.Next
		head.Next = tmpNode
	}
	return head.Next
}

type Stack []int

func (stack Stack) Len() int {
	return len(stack)
}

func (stack *Stack) Push(value int)  {
	*stack = append(*stack, value)
}

func (stack *Stack) Pop() int  {
	if len(*stack) == 0 {
		return 0
	}
	tmp := *stack
	value := tmp[len(*stack) - 1]
	*stack = tmp[: len(*stack) - 1]
	return value
}

func main() {
	l1 := &ListNode{
		Val:  7,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  4,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	l2 := &ListNode{
		Val:  5,
		Next: &ListNode{
			Val:  6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	addTwoNumbers(l1, l2)
}
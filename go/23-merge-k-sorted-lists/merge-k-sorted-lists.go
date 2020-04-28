package main

func main() {
	a:= []*ListNode{&ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	},
	&ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	},
	}

	mergeKLists(a)
}
 type ListNode struct {
     Val int
     Next *ListNode
 }


 func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	head := ListNode{
		Val:  0,
		Next: lists[0],
	}

	for i := 1 ; i < len(lists); i ++ {
		tmpHead := head.Next
		tmpList := lists[i]
		newList := &ListNode{
			Val:  0,
			Next: nil,
		}
		tmpNewList := newList
		for tmpHead != nil && tmpList != nil {
			if tmpHead.Val < tmpList.Val {
				tmpNewList.Next = &ListNode{
					Val:  tmpHead.Val,
					Next: nil,
				}
				tmpNewList = tmpNewList.Next
				tmpHead = tmpHead.Next
			} else {
				tmpNewList.Next = &ListNode{
					Val:  tmpList.Val,
					Next: nil,
				}
				tmpNewList = tmpNewList.Next
				tmpList = tmpList.Next
			}
		}
		if tmpHead == nil {
			tmpNewList.Next = tmpList
		}else {
			tmpNewList.Next = tmpHead
		}
		head.Next = newList.Next
	}
	return head.Next
}
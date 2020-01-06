<?php

 class ListNode {
     public $val = 0;
     public $next = null;
     function __construct($val) { $this->val = $val; }
 }

 class Solution {

    /**
     * @param ListNode $head
     * @param Integer $m
     * @param Integer $n
     * @return ListNode
     */
    function reverseBetween($head, $m, $n) {
        if ($m === 0) {
            $m++;
            $n++;
        }
        $i = 0;
        $tmp = new ListNode(null);
        $tmp->next = $head;
        $p1 = &$tmp;
        $p2 = &$tmp->next;
        while (true) {
            if ($i < $m - 1) {
                $p1 = &$p1->next;
                $p2 = &$p2->next;
            }elseif ($m - 1 <= $i && $i< $n -1 ) {
                $target = new ListNode($p2->next->val);
                $target->next = $p1->next;
                $p1->next = $target;
                if ($i == $m-1) {
                    $p2 = &$p2->next;
                }
                $p2->next = $p2->next->next;
            } else {
                break;
            }
            $i++;
        }
        return $tmp->next;
    }
}

$list = new ListNode(1);
$list->next = new ListNode(2);
$list->next->next = new ListNode(3);
$list->next->next->next = new ListNode(4);
$list->next->next->next->next = new ListNode(5);

$a = new Solution();
print_r($a->reverseBetween($list, 0,5));

/**
 *  解题思路：
 *    双指针遍历链表 快慢指针相差一个节点  遇到开始反转位置的前一个节点 慢指针停止  将快指针节点交给慢指针节点的next
 *    当快指针是慢指针的next时候 需要将快指针移动到next  因为慢指针新增node时 快指针引用 也会多一个node
 *    时间复杂度O(n)
 */
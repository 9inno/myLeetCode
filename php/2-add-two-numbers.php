<?php


class ListNode {
    public $val = 0;
    public $next = null;
    function __construct($val) { $this->val = $val; }
}

class Solution {

    /**
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    function addTwoNumbers($l1, $l2) {
        $result = new ListNode(null);
        $tmp = &$result;
        while (true) {

            $sum = $l1->val + $l2->val;
            if ($sum >= 10) {
                $sum = $sum % 10;
                $tmp->val = $sum;
                $l1->next->val++;
            }else {
                $tmp->val = $sum;
            }
            if ($l1->next === null && $l2->next === null){
                break;
            }
            $tmp->next = new ListNode(null);
            $l1 = $l1->next;
            $l2 = $l2->next;
            $tmp = &$tmp->next;
        }
        return $result;
    }
}

/**
 * 解题思路：
 * 同时遍历两个链表，相加节点的值，
 * 如果和大于10，则对10取模，结果为新链表节点的值，
 * 后随意向两原始链表任意一条下一节点值 +1
 * 如果不小于10 则直接将和赋予新链表节点的值，
 * 直到两链表节点 next ===null 停止
 * 时间复杂度 O(max(m,n))
 *
 */
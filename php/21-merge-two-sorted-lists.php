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
    function mergeTwoLists($l1, $l2) {
        $result = new ListNode(0);
        $tmp = &$result;
        while (true) {
            if (($l1->val < $l2->val && $l1!== null) || $l2 === null) {
                $tmp->val = $l1->val;
                $l1 = $l1->next;
            }else {
                $tmp->val = $l2->val;
                $l2 = $l2->next;
            }
            if ($l1 === null && $l2 === null) {
                break;
            }
            $tmp->next = new ListNode(null);
            $tmp = &$tmp->next;
        }
        return $result;
    }
}

/**
 * 解题思路：
 * 同时遍历两链表，
 * 判断当前节点值的大小，
 * 小的值放入新的链表，并指向next 继续判断
 * 时间复杂度O(max(m,n))
 */
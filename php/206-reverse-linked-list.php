<?php
 class ListNode {
     public $val = 0;
     public $next = null;
     function __construct($val) { $this->val = $val; }
 }

class Solution {

    /**
     * @param ListNode $head
     * @return ListNode
     */
    function reverseList($head) {
        $result = new ListNode(null);
        while ($head != null) {
            $tmp = new ListNode($head->val);
            $tmp->next = $result->next;
            $result->next = $tmp;
            $head = $head->next;
        }

        return $result->next;
    }
}

/**
 * 解题思路：
 *   遍历链表 将节点插入新链表的开头
 *   时间复杂度O(n)
 */
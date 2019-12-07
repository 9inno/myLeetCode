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
    function deleteDuplicates($head) {
        if ($head->next === null) {
            return $head;
        }
        while ($head->val === $head->next->val) {
            $head = $head->next;
        }
        $head->next = $this->deleteDuplicates($head->next);
        return $head;
    }
}

/**
 * 解题思路：
 *  递归 对比当前节点与下一节点 相同则使当前节点指向下一节点的下一节点
 *  时间复杂度O(n)
 *
 */
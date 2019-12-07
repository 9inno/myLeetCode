<?php

class ListNode {
    public $val = 0;
    public $next = null;
    function __construct($val) { $this->val = $val; }
}

class Solution {

    /**
     * @param ListNode $head
     * @param Integer $n
     * @return ListNode
     */
    function removeNthFromEnd($head, $n) {
        $i = 0;
        $fast = $head;
        $slow = null;
        while (true) {
            if ($i == $n) {
                $slow = &$head;
            }
            $i++;
            if ($fast->next === null) {
                break;
            }
            $fast = $fast->next;
            if ($slow) {
                $slow = &$slow->next;
            }

        }
        if (isset($slow->next->next)){
            $slow->next = $slow->next->next;
        }else {

            if (!isset($slow)) {
                $head = $head->next;
            } else {
                $slow->next =null;
            }
        }
        return $head;
    }
}

/**
 * 解题思路：
 * 采用双指针遍历链表
 * 双指针间隔 $n 个节点
 * 当快指针达到链表结尾时 从慢指针切断链表
 * 时间复杂度O(n)
 */
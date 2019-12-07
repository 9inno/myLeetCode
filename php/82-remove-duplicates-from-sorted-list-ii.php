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
        $arr =[];
        $tmp = new ListNode(0);
        $tmp->next = &$head;
        while ($tmp->next !== null) {
            if ($tmp->next->val === $tmp->next->next->val) {
                $arr[$tmp->next->val] = true;
            }
            if (isset($arr[$tmp->next->val])) {
                $tmp->next = $tmp->next->next;
            }else {
                $tmp = &$tmp->next;
            }
        }
        return $head;
    }
}

/**
 * 解题思路：
 * 新开数组储存重复的节点的值
 * 遍历链表 若当前节点在数组内存在值 则删除当前节点
 * 需要空的头节点 控制边界值 即第一个节点与第二个节点相同值
 * 时间复杂度O(n)
 */
<?php

class ListNode {
    public $val = 0;
    public $next = null;
    function __construct($val) { $this->val = $val; }
}

class Solution {

    /**
     * @param ListNode $head
     * @param Integer $val
     * @return ListNode
     */
    function removeElements($head, $val) {
        $emptyHead = new ListNode(0);
        $emptyHead->next = &$head;
        $tmp = &$emptyHead;
        while ($tmp->next !== null) {
           if ($tmp->next->val !== $val) {
               $tmp = &$tmp->next;
               continue;
           } else {
               $tmp->next = $tmp->next->next;
           }
        }
        return $head;
    }
}

//test
$a = new ListNode(2);
$b = & $a;
for ($i = 0; $i < rand(1,20); $i++) {
    $b->next = new ListNode(rand(1,20));
    $b = &$b->next;
}
print_r($a);

$s = new Solution();
print_r($s->removeElements($a, 2));

/**
 * 解题思路：
 *   创建空头节点
 *   从空节点遍历下一节点 节点值等于目标 则删除
 *   时间复杂度O(n)
 */
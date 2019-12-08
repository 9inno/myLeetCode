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
    function swapPairs($head) {

        if ($head === null || $head->next ===null) {
            return $head;
        }

        $tmp = $head->next;
        $head->next = $this->swapPairs($tmp->next);
        $tmp->next = $head;
        return  $tmp;
    }
}

//test
$a = new ListNode(1);
$b = & $a;
for ($i = 0; $i < rand(1,20); $i++) {
    $b->next = new ListNode(rand(1,20));
    $b = &$b->next;
}
print_r($a);

$s = new Solution();
print_r($s->swapPairs($a));


/**
 * 解题思路：
 *  当前节点 与后一节点 互换
 *  下一节点的下一节点丢进递归
 *  注意先丢入递归再互换节点
 */
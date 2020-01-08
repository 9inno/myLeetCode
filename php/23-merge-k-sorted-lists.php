<?php

class ListNode {
    public $val = 0;
    public $next = null;
    function __construct($val) { $this->val = $val; }
}

class Solution {

    /**
     * @param ListNode[] $lists
     * @return ListNode
     */
    function mergeKLists($lists) {
        $result = new ListNode(null);

        foreach ($lists as $list) {
            $p = &$result;
            while ($list != null) {
                if ($p->next === null) {
                    $p->next = $list;
                    break;
                }
                if ($list->val < $p->next->val) {
                    $tmp = $p->next;
                    $p->next = clone $list;
                    $p->next->next = $tmp;
                    $p = &$p->next;
                    $list = $list->next;
                }else {
                    $p = &$p->next;
                }
            }
        }

        return $result->next;
    }
}

$list1 = new ListNode(1);
$list1->next = new ListNode(4);
$list1->next->next = new ListNode(5);

$list2 = new ListNode(1);
$list2->next = new ListNode(3);
$list2->next->next = new ListNode(4);

$list3 = new ListNode(2);
$list3->next = new ListNode(6);

$a = new Solution();
print_r($a->mergeKLists([]));


/**
 *  解题思路：
 *   定义一条空链表 逐一与数组中的链表 遍历比较合并
 *   时间复杂度O(kN)
 */
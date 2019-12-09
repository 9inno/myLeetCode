<?php

class ListNode {
    public $val = 0;
    public $next = null;
    function __construct($val) { $this->val = $val; }
}

class Solution {

    /**
     * @param ListNode $head
     * @param Integer $x
     * @return ListNode
     */
    function partition($head, $x)
    {
        $emptyHead = new ListNode(0);
        $emptyHead->next = $head;
        $tmp = &$emptyHead;
        while ($tmp->next !== null) {
            if ($tmp->next->val >= $x) {
                break;
            }
            $tmp = & $tmp->next;
        }
        $lowHalf = new ListNode(0);
        $lowHalf->next =  $tmp->next;
        $lowHalfTmp = & $lowHalf;
        while (true) {
            if (isset($lowHalfTmp->next->val) && $lowHalfTmp->next->val < $x) {
                $tmp->next = $lowHalfTmp->next;
                $lowHalfTmp->next = $lowHalfTmp->next->next;
                $tmp = &$tmp->next;
            }else{
                $lowHalfTmp = & $lowHalfTmp->next;
            }


            if ($lowHalfTmp->next === null) {
                break;
            }

        }
        $tmp->next = $lowHalf->next ?? null;
        return $emptyHead->next;
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
print_r($s->partition($a, 2));

/**
 * 解题思路：
 *   再第一个大于目标值的节点切断 链表
 *   遍历后半部分链表 小于目标值的节点依此 接到前半部分链表末尾
 *   遍历结束后 拼接前半部分和后半部分链表
 *   时间复杂度O(n)
 */
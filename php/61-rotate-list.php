<?php

class ListNode {
    public $val = 0;
    public $next = null;
    function __construct($val) { $this->val = $val; }
}

class Solution {

    /**
     * @param ListNode $head
     * @param Integer $k
     * @return ListNode
     */
    function rotateRight($head, $k) {
        if ($k === 0) {
            return $head;
        }
        $count = 0;
        $tmp = [];
        while ($head !== null) {
            $tmp[] = $head->val;
            $head = $head->next;
            $count++;
        }
        if ($count === 0) {
            return null;
        }

        $result = new ListNode(null);
        $node = &$result->next;
        $i = $count - ($k % $count) - 1;
        $j = $i + 1;
        while (true) {
            if ($j === $count) {
                $j = 0;
            }
            $node = new ListNode($tmp[$j]);
            $node = &$node->next;
            if ($j === $i ) {
                break;
            }
            $j++;
        }

        return $result->next;

    }
}

$list = new ListNode(1);
$list->next = new ListNode(2);
$list->next->next = new ListNode(3);
$list->next->next->next = new ListNode(4);
$list->next->next->next->next = new ListNode(5);

$a = new Solution();
print_r($a->rotateRight($list, 1));


/**
 *  解题思路：
 *   向右移动链表节点 实际上是找 倒数 k 个 节点 切断
 *   前面部分拼接到后边部分
 *   先遍历链表 存入数组 确定链表长度
 *   倒数低k个 用 count  - (k % count) 计算 然后从该位置遍历数组生成新的链表
 *   指针达到数组最后时 将指针重新定位到数组开头
 *   当指针回到最开始的位置 结束遍历
 *   时间复杂度O(n)
 */
<?php

class TreeNode {
    public $val = null;
    public $left = null;
    public $right = null;
    function __construct($value) { $this->val = $value; }
}

class Solution {

    /**
     * @param TreeNode $root
     * @return Integer[]
     */
    function largestValues($root) {
        $queue = [];
        $result = [];
        if ($root !== null) {
            $queue[] = $root;
        }
        while (!empty($queue)) {
            $tmp = [];
            $max = null;
            foreach ($queue as $node) {
                if ($max === null) {
                    $max = $node->val;
                } else {
                    $max = max($max, $node->val);
                }
                $node->left !== null && $tmp[] = $node->left;
                $node->right !== null && $tmp[] = $node->right;
            }
            $result[] = $max;
            $queue = $tmp;
        }
        return $result;
    }
}

/**
 * 解题思路 ：
 *   按层遍历树 用数组存储下一层节点  并计算本层节点最大值
 *   时间复杂度O(n)
 */
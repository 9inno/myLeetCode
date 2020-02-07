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
     * @return Integer
     */
    function findBottomLeftValue($root) {
        $queue = [$root];
        while (!empty($queue)) {
            $tmp = [];
            foreach ($queue as $node) {
                /**@var TreeNode $node**/
                $node->left !== null && $tmp[] = $node->left;
                $node->right !== null && $tmp[] = $node->right;
            }

            if (empty($tmp)) {
                return $queue[0]->val;
            } else {
                $queue = $tmp;
            }
        }

    }
}

/**
 * 解题思路:
 *    按层遍历树 广度优先搜索
 *    时间复杂度O(n)
 */
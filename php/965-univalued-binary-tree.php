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
     * @return Boolean
     */
    function isUnivalTree($root) {
        $queue = [];
        if ($root !== null) {
            $queue[] = $root;
            $value = $root->val;
        }
        while (!empty($queue)) {
            $tmp = [];
            foreach ($queue as $node) {
                /**@var TreeNode $node**/
                if (isset($value) && $node->val !== $value) {
                    return false;
                }
                $node->left !== null && $tmp[] = $node->left;
                $node->right !== null && $tmp[] = $node->right;
            }
            $queue = $tmp;
        }

        return true;
    }
}

/**
 * 解题思路
 *   按层遍历树 广度优先遍历
 *   时间复杂度O(n)
 */
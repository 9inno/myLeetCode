<?php

class TreeNode {
    public $val = null;
    public $left = null;
    public $right = null;
    function __construct($value) { $this->val = $value; }
}

class Solution {
    private $result = 0;
    /**
     * @param TreeNode $root
     * @return Integer
     */
    function findTilt($root) {

        $this->recursion($root);
        return $this->result;
    }

    function recursion($root) {
        if ($root === null) {
            return 0;
        }

        $left = $this->recursion($root->left  );
        $right = $this->recursion($root->right);

        $sum = $left + $right;
        $this->result += abs($left - $right);
        return $sum + $root->val;

    }
}

/**
 * 解题思路：
 *   递归遍历树 每个节点返回 子节点的和 同时计算坡度
 *   时间复杂度O(n)
 */
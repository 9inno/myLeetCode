<?php

class TreeNode {
    public $val = null;
    public $left = null;
    public $right = null;
    function __construct($value) { $this->val = $value; }
}

class Solution {
    private $result = [];
    /**
     * @param TreeNode $root
     * @return String[]
     */
    function binaryTreePaths($root) {
        if ($root !== null) {
            $this->recursion($root, []);
        }
        return $this->result;
    }

    function recursion($root, $str) {
        if ($root === null) {
            return ;
        }
        if ($root->left === null && $root->right === null) {
            $this->result [] = implode('->', array_merge($str, [$root->val]));
        } else{
            $this->recursion($root->left,  array_merge($str, [$root->val]));
            $this->recursion($root->right,  array_merge($str, [$root->val]));
        }
        return ;
    }
}


/**
 * 解题思路：
 *   与129题相同，
 *   深度优先遍历树 放入数组 遍历到叶子节点 计入结果
 *   时间复杂度O(n)
 */
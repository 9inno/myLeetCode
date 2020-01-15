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
     * @return Integer[]
     */
    function preorderTraversal($root) {

        if ($root  === null) {
            return [];
        }
        $this->result[] = $root->val;
        $this->preorderTraversal($root->left);
        $this->preorderTraversal($root->right);
        return $this->result;

    }

    function preorderTraversal2($root) {
        if ($root === null) {
            return [];
        }
        $result = [];
        $stack = [$root];
        while (!empty($stack)) {
            $node = array_pop($stack);
            $result[] = $node->val;
            $node->right !== null && array_push($stack, $node->right);
            $node->left !== null && array_push($stack, $node->left);
        }

        return $result;

    }
}

/**
 *
 *  解题思路：
 *   1. 前序遍历二叉树 递归方法
 *   2. 迭代 用数组模拟栈 遍历节点
 * 
 */
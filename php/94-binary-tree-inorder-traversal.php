<?php

class TreeNode {
    public $val = null;
    public $left = null;
    public $right = null;
    function __construct($value) { $this->val = $value; }
}

class Solution {

    public $result = [];
    /**
     * @param TreeNode $root
     * @return Integer[]
     */
    function inorderTraversal($root) {
        if ($root === null ){
            return [];
        }

        $this->inorderTraversal($root->left);
        array_push($this->result, $root->val);
        $this->inorderTraversal($root->right);
        return $this->result;
    }
}

$tree = new TreeNode(5);
$tree->right = new TreeNode(6);
$tree->left = new TreeNode(7);
$tree->left->left = new TreeNode(10);
$tree->right->right = new TreeNode(11);
$tree->right->left = new TreeNode(15);
$tree->left->right = new TreeNode(20);


$a = new Solution();
print_r($a->inorderTraversal($tree));

/**
 * 解题思路：
 *   中序遍历二叉树 遍历顺序 左子节点 -> 父节点 -> 右子节点
 *   时间复杂度 O(n)
 */
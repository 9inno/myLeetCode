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
     * @return TreeNode
     */
    function invertTree($root) {

        if ($root === null) {
            return ;
        }
        list($root->left,$root->right) = [$root->right, $root->left];
        $this->invertTree($root->left);
        $this->invertTree($root->right);
        return $root;
    }
}


$tree = new TreeNode(5);
$tree->left = new TreeNode(9);
$tree->right = new  TreeNode(11);
$tree->right->left = new  TreeNode(15);
$tree->right->right = new TreeNode(17);
$a = new Solution();
print_r($a->invertTree($tree));



/**
 * 解题思路：
 *   递归交换左右树
 *   时间复杂度O(n)
 */
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
     * @param Integer $target
     * @return mixed
     */
    function removeLeafNodes($root, $target) {

        if($root === null) {
            return ;
        }
        if ($root->right === null && $root->left === null && $root->val == $target) {
            return "flag";
        }

        $left = $this->removeLeafNodes($root->left, $target);
        $right = $this->removeLeafNodes($root->right, $target);
        $left ==='flag' && $root->left =null;
        $right ==='flag' && $root->right =null;

        if ($root->right === null && $root->left === null && $root->val == $target) {
            return "flag";
        }
        return $root;

    }
}

$tree = new TreeNode(1);
$tree->left = new TreeNode(2);
$tree->left->left = new TreeNode(2);
$tree->right = new TreeNode(3);
$tree->right->left =new TreeNode(2);
$tree->right->right = new TreeNode(4);

$a = new Solution();
print_r($a->removeLeafNodes($tree, 2));
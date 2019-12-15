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
     * @param Integer $sum
     * @return Boolean
     */
    function hasPathSum($root, $sum) {

        if ($root === null) {
            return false;
        }

        $sum -= $root->val;
        if ($root->left === null && $root->right === null) {
            return $sum == 0;
        }

        return $this->hasPathSum($root->left, $sum) || $this->hasPathSum($root->right, $sum);
    }

}

$tree = new TreeNode(5);
$tree->left = new TreeNode(9);
$tree->right = new  TreeNode(11);
$tree->right->left = new  TreeNode(15);
$tree->right->right = new TreeNode(17);

$a = new Solution();
var_dump($a->hasPathSum($tree,31));


/**
 * 解题思路：
 *   目标值减去每个节点的值  进入递归到叶子节点  判断差是否为0
 *   时间复杂度 O(n)
 */
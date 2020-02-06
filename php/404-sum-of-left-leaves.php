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
    function sumOfLeftLeaves($root) {
        $this->recursion($root);
        return $this->result;
    }

    function recursion($root) {
        if ($root === null) {
            return 0;
        }
        if ($root->left === null && $root->right === null) {
            return $root->val;
        }
        $left = $this->recursion($root->left);
        $right = $this->recursion($root->right);
        $this->result += $left;
    }
}

$tree = new TreeNode(1);
$tree->left = new TreeNode(1);
$tree->left->left = new TreeNode(1);
$tree->right = new TreeNode(3);
$tree->right->left = new TreeNode(5);
$a = new Solution();
echo $a->sumOfLeftLeaves($tree);
/**
 * 解题思路：
 *   只记录左叶子节点 的值  遍历全树
 * 时间复杂度O(n)
 */
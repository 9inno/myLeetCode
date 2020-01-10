<?php

class TreeNode {
    public $val = null;
    public $left = null;
    public $right = null;
    function __construct($value) { $this->val = $value; }
}

class Solution {

    private $count = 0;
    /**
     * @param TreeNode $root
     * @return Integer
     */
    function countNodes($root) {

        if ($root === null) {
            return 0;
        }
        $this->count ++;
        $this->countNodes($root->left);
        $this->countNodes($root->right);
        return $this->count;
    }
}

$tree = new TreeNode(5);
$tree->left = new TreeNode(9);
$tree->right = new  TreeNode(11);
$tree->right->left = new  TreeNode(15);
$tree->right->right = new TreeNode(17);

$a = new Solution();
echo $a->countNodes($tree);

/**
 * 解题思路：
 *    遍历二叉树 记录节点个数
 *    时间复杂度O(n)
 */
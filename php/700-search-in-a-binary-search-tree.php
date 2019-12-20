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
     * @param Integer $val
     * @return TreeNode
     */
    function searchBST($root, $val) {

        if ($root === null) {
            return null;
        }
        if ($root->val === $val) {
            return $root;
        }
        $left = $this->searchBST($root->left,$val);
        $right = $this->searchBST($root->right,$val);
        return $left === null ? $right : $left;
    }
}


//test
$tree = new TreeNode(5);
$tree->left = new TreeNode(9);
$tree->right = new  TreeNode(11);
$tree->right->left = new  TreeNode(15);
$tree->right->right = new TreeNode(17);
$tree->left->left = new  TreeNode(16);
$a = new Solution();
print_r($a->searchBST($tree,9));


/**
 *  解题思路：
 *   深度优先搜索 搜索到则返回当前节点  搜索不到则返回null
 *   时间复杂度O(n)
 */
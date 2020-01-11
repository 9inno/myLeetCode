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
     * @return Integer[]
     */
    function rightSideView($root) {
        $arr = [];
        $result = [];
        if ($root !== null) {
            $arr[] = $root;
        }
        while (!empty($arr)) {
            $tmp = [];
            foreach ($arr as $node) {
                $node->left !== null && $tmp[] = $node->left;
                $node->right !== null && $tmp[] = $node->right;
            }
            $result[] = end($arr)->val;
            $arr = $tmp;
        }
        return  $result;
    }
}

$tree = new TreeNode(5);
$tree->left = new TreeNode(9);
$tree->right = new  TreeNode(11);
$tree->right->left = new  TreeNode(15);
$tree->right->right = new TreeNode(17);

$a = new Solution();
print_r($a->rightSideView($tree));

/**
 * 解题思路：
 *   用数组记录树每层的节点 取每层最右边的树节点值
 *   时间复杂度O(n)
 */
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
     * @return NULL
     */
    function flatten(&$root) {

        $stack = [];
        array_push($stack,$root);
        $link = new TreeNode(null);
        $tmp = $link;
        while (!empty($stack)) {
            $node = array_pop($stack);
            $tmp->right = new TreeNode($node->val);

            if ($node->right !== null) {
                array_push($stack,$node->right);
            }

            if ($node->left !== null) {
                array_push($stack,$node->left);
            }

            $tmp = $tmp->right;
        }

        $root = $link->right;

    }


}


$tree = new TreeNode(1);
$tree->left = new TreeNode(2);
$tree->right = new TreeNode(5);
$tree->left->left = new TreeNode(3);
$tree->left->right = new TreeNode(4);
$tree->right->right = new TreeNode(6);


$a = new Solution();
$a->flatten($tree);
print_r($tree);


/**
 *  解题思路：
 *   前序遍历  递归超时 改用非递归 用数组模拟栈降节点按先右后左的顺序压入栈
 *   时间复杂度O(n)
 *   TODO 原地展开
 */
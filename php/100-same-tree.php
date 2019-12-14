<?php

class TreeNode {
    public $val = null;
    public $left = null;
    public $right = null;
    function __construct($value) { $this->val = $value; }
}

class Solution {

    /**
     * @param TreeNode $p
     * @param TreeNode $q
     * @return Boolean
     */
    function isSameTree($p, $q) {
        if ($p === null && $q ===null) {
            return true;
        }

        if ($p->val !== $q->val) {
            return false;
        }

        $leftSame = $this->isSameTree($p->left, $q->left);
        $rightSame = $this->isSameTree($p->right, $q->right);

        return $leftSame && $rightSame;
    }
}

$treeA = new TreeNode(5);
$treeA->left = new TreeNode(9);
$treeA->right = new  TreeNode(11);
$treeA->right->left = new  TreeNode(15);
$treeA->right->right = new TreeNode(17);

$treeB = new TreeNode(5);
$treeB->left = new TreeNode(9);
$treeB->right = new  TreeNode(11);
$treeB->right->left = new  TreeNode(15);
$treeB->right->right = new TreeNode(17);

$a = new Solution();
var_dump($a->isSameTree($treeA,$treeB));


/**
 * 解题思路 ：
 *   同时 深度优先遍历两颗树 判断当前节点值是否相等
 */
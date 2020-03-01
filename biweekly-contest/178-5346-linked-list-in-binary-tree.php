<?php

class ListNode {

    public $val = 0;
    public $next = null;
    function __construct($val) { $this->val = $val; }
}


class TreeNode {
    public $val = null;
    public $left = null;
    public $right = null;
    function __construct($value) { $this->val = $value; }
}

class Solution {
    /**
     * @param ListNode $head
     * @param TreeNode $root
     * @return Boolean
     */

    function isSubPath($head, $root) {
        if ($head == null) {
            return true;
        }
        if ($root == null) {
            return false;
        }

        return $this->recursion($head, $root) || $this->isSubPath($head, $root->left) || $this->isSubPath($head, $root->right);
    }

    function recursion($head, $root) {
        if ($head == null) {
            return true;
        }
        if ($root == null) {
            return false;
        }
        if ($head->val != $root->val) {
            return false;
        }

        return $this->recursion($head->next,$root->left) || $this->recursion($head->next, $root->right);
    }


}

$tree = new TreeNode(1);
$tree->right = new TreeNode(1);
$tree->right->left = new TreeNode(10);
$tree->right->right = new TreeNode(1);
$tree->right->left->left = new TreeNode(9);

$list = new ListNode(1);
$list->next = new ListNode(10);

$a = new Solution();
var_dump($a->isSubPath($list, $tree));
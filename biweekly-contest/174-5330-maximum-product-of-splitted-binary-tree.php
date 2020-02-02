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
    function maxProduct($root) {
        $sum = $this->sum($root);
        $this->mul($root, $sum);
        return $this->result % (10**9 + 7);
    }

    function sum($root) {
        if ($root == null) {
            return 0;
        }
        $left = $this->sum($root->left);
        $right = $this->sum($root->right);
        return $left + $right + $root->val;
    }
    function mul($root, $sum) {
        if ($root == null) {
            return 0;
        }
        $left = $this->mul($root->left, $sum);
        $right = $this->mul($root->right, $sum);
        $tmp = $left + $right + $root->val;
        $target = $sum - $tmp;
        $this->result = max($this->result, $target * $tmp);
        return $tmp;
    }
}


$tree = new TreeNode(1);
$tree->left = new TreeNode(2);
$tree->right = new TreeNode(3);
$tree->left->left = new TreeNode(4);
$tree->left->right = new TreeNode(5);
$tree->right->left = new TreeNode(6);

$a = new Solution();
echo $a->maxProduct($tree);
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
    function sumNumbers($root) {

        if ($root !== null) {
            $this->recursion($root, 0);
        }
        return $this->result;
    }

    function recursion($root, $str) {
        if ($root === null) {
            return ;
        }
        if ($root->left === null && $root->right === null) {

            $this->result += $str * 10 + $root->val;
        } else{
            $this->recursion($root->left,  $str * 10 + $root->val);
            $this->recursion($root->right,  $str * 10 + $root->val);
        }
        return ;
    }
}


$tree = new TreeNode(1);
$tree->left = new TreeNode(2);
$tree->right = new TreeNode(3);

$a = new Solution();
echo $a->sumNumbers($tree);

/**
 * 解题思路：
 *   递归 深度优先遍历 叶子节点为没有子节点的节点 所以在left right 同事为null 时候 做加和
 *   时间复杂度O(n)
 */
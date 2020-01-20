<?php

class TreeNode {
    public $val = null;
    public $left = null;
    public $right = null;
    function __construct($value) { $this->val = $value; }
}

class Solution {
    private $result = [];
    /**
     * @param TreeNode $root
     * @return Integer[]
     */
    function postorderTraversal($root) {

        if($root === null) {
            return [];
        }

        $this->postorderTraversal($root->left);
        $this->postorderTraversal($root->right);
        $this->result[] = $root->val;
        return $this->result;
    }
}
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
     * @return Integer
     */
    function sumEvenGrandparent($root) {
        $queue = [];
        $result = 0;
        if ($root !== null) {
            $queue[] = $root;
        }

        while (!empty($queue)) {
            $tmp = [];
            foreach ($queue as $node) {
                $node->left !== null && $tmp[]=$node->left;
                $node->right !== null && $tmp[]=$node->right;
                if ($node->val % 2 == 0) {
                    $node->left !== null && $node->left->left !== null && $result+= $node->left->left->val;
                    $node->left !== null && $node->left->right !== null && $result+= $node->left->right->val;
                    $node->right !== null && $node->right->left !== null && $result+= $node->right->left->val;
                    $node->right !== null && $node->right->right !== null && $result+= $node->right->right->val;
                }
            }
            $queue = $tmp;
        }
        return $result;

    }
}
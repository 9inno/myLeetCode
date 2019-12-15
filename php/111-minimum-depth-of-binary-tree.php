<?php

/**
 * Definition for a binary tree node.
 * Class TreeNode
 */
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
    function minDepth($root) {
        if ($root === null) {
            return 0;
        }

        $leftHeight = $this->minDepth($root->left);
        $rightHeight = $this->minDepth($root->right);
        if ($root->left === null || $root->right === null) {
            return $leftHeight + $rightHeight + 1 ;
        }

        return min($leftHeight,$rightHeight) + 1;
    }

}

$tree = new TreeNode(5);
$tree->left = new TreeNode(9);
$tree->right = new  TreeNode(11);
$tree->right->left = new  TreeNode(15);
$tree->right->right = new TreeNode(17);

$a = new Solution();
echo $a->minDepth($tree);

/**
 * 解题思路：
 *   与104题思路一样  递归 去左右子树深度小的+1
 *   需要注意 只有左树或只有右树的情况 会产生深度为0的比较
 */
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
     * @return Integer[][]
     */
    function levelOrder($root) {

        $queue = [$root];
        $result = [];

        while (!empty($queue)) {
            $tmp = [];
            $resultTmp = [];
            foreach ($queue as $treeNode) {
                $treeNode->left !== null && $tmp[] = $treeNode->left;
                $treeNode->right !== null && $tmp[] = $treeNode->right;
                $treeNode->val !== null && $resultTmp[] = $treeNode->val;
            }
            !empty($resultTmp) && $result[] = $resultTmp;
            $queue = $tmp;
        }
        return $result;
    }
}

$tree = new TreeNode(5);
$tree->left = new TreeNode(9);
$tree->right = new  TreeNode(11);
$tree->right->left = new  TreeNode(15);
$tree->right->right = new TreeNode(17);

$a = new Solution();
print_r($a->levelOrder(null));

/**
 * 解题思路：
 *   一层一层遍历 用数组暂存下一层的节点
 *   时间复杂度O(n)
 */
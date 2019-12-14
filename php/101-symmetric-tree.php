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
     * @return Boolean
     */
    function isSymmetric($root) {
        if ($root === null) {
            return true;
        }

        $queue = [$root];
        while (!empty($queue)) {
            $tmp = [];
            $tmpQueue = [];
            foreach ($queue as $treeNode) {
                $tmp[] = $treeNode->val;
                if ($treeNode !== null) {
                    $tmpQueue[] = $treeNode->left;
                    $tmpQueue[] = $treeNode->right;
                }
            }
            $count = count($tmp);
            for ($i = 0; $i < $count; $i ++) {
                if ($tmp[$i] !== $tmp[$count - $i - 1]) {
                    return false;
                }
            }
            $queue = $tmpQueue;
        }

        return true;
    }
}
//test
$treeA = new TreeNode(5);
$treeA->left = new TreeNode(9);
$treeA->right = new  TreeNode(9);
$treeA->right->left = new  TreeNode(15);
$treeA->left->right = new  TreeNode(15);
//$treeA->right->right = new TreeNode(17);

$a = new Solution();
var_dump($a->isSymmetric($treeA));

/**
 * 解题思路：
 *   广度优先遍历  每层节点值放入临时数组  遍历数组是否对称
 *   TODO 可优化： 一次遍历 O(n) 复杂度
 */
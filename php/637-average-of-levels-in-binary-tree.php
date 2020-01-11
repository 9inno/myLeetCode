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
     * @return Float[]
     */
    function averageOfLevels($root) {
        if ($root === null) {
            return [];
        }
        $arr = [$root];
        $result = [];
        while (!empty($arr)) {
            $tmp = [];
            $resultTmp = 0;
            foreach ($arr as $node) {
                $resultTmp += $node->val;
                $node->left !== null && $tmp[] = $node->left;
                $node->right !== null && $tmp[] = $node->right;
            }
            $result[] = $resultTmp / count($arr);
            $arr = $tmp;
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
print_r($a->averageOfLevels($tree));

/**
 *  解题思路：
 *    按层遍历树  记录每层树节点的平均值
 *    时间复杂度O(n)
 */
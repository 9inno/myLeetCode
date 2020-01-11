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
    function levelOrderBottom($root) {

        if ($root === null) {
            return [];
        }
        $arr = [$root];
        $result = [];
        while (!empty($arr)) {
            $tmp = [];
            $resultTmp = [];
            foreach ($arr as $node) {
                $resultTmp[] = $node->val;
                $node->left !== null && $tmp[] = $node->left;
                $node->right !== null && $tmp[] = $node->right;
            }

            array_unshift($result,$resultTmp);
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
print_r($a->levelOrderBottom(null));

/**
 * 解题思路：
 *   用 数组记录下一层的树节点  按层遍历树
 *   时间复杂度O(n)
 */
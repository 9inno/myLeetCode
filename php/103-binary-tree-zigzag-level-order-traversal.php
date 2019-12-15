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
    function zigzagLevelOrder($root) {
        $result = [];
        $queue = [$root];
        $n = 1;
        while (!empty($queue)) {
            $count = count($queue);
            $mod = $n % 2;
            $tmp = [];
            $queueTmp = [];
            if ($mod == 1) {
                for ($i = 0; $i < $count; $i++) {
                    $queue[$i]->val !== null && $tmp[] = $queue[$i]->val;
                    $queue[$i]->left !== null && $queueTmp[] = $queue[$i]->left;
                    $queue[$i]->right !== null && $queueTmp[] = $queue[$i]->right;
                }
            } else {
                for ($i = $count - 1; $i >= 0 ; $i--) {
                    $queue[$i]->val !== null && $tmp[] = $queue[$i]->val;
                    $reverse =[];
                    $queue[$i]->left !== null && $reverse[] = $queue[$i]->left;
                    $queue[$i]->right !== null && $reverse[] = $queue[$i]->right;
                    $queueTmp = array_merge($reverse,$queueTmp);
                }
            }
            !empty($tmp) && $result[] = $tmp;
            $queue = $queueTmp;
            $n++;
        }
        return $result;
    }
}

//test
$tree = new TreeNode(5);
$tree->left = new TreeNode(9);
$tree->right = new  TreeNode(11);
$tree->right->left = new  TreeNode(15);
$tree->right->right = new TreeNode(17);
$tree->left->left = new  TreeNode(16);
$tree->left->right = new TreeNode(18);
$a = new Solution();
print_r($a->zigzagLevelOrder($tree));



/**
 * 解题思路：
 *   按层遍历  按层数对2取模 确定正向逆向
 *   逆向遍历时 需要注意处理下一层的顺序
 *   时间复杂度O(n)
 */
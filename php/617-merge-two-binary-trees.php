<?php

class TreeNode {
    public $val = null;
    public $left = null;
    public $right = null;
    function __construct($value) { $this->val = $value; }
}

class Solution {

    /**
     * @param TreeNode $t1
     * @param TreeNode $t2
     * @return TreeNode
     */
    function mergeTrees($t1, $t2) {

        if ($t1 === null && $t2 === null) {
            return ;
        }
        $t1->val += $t2->val ?? 0;

        if ($t1->left === null && $t2->left !== null) {
            $t1->left = new TreeNode(0);
        }
        if ($t1->right === null && $t2->right !== null) {
            $t1->right = new TreeNode(0);
        }
        $this->mergeTrees($t1->left, $t2->left ?? null);
        $this->mergeTrees($t1->right , $t2->right ?? null);
        return $t1;
    }
}


/**
 * 解题思路：
 *   同时遍历两棵树 以一棵树为基准  将另外一棵树同节点向上添加 若基准数不存在节点而另外树存在节点则新建树节点
 *   时间复杂度 O(max(n,m))
 */
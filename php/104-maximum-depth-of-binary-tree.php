<?php

class TreeNode {
    public $val = null;
    public $left = null;
    public $right = null;
    function __construct($value) { $this->val = $value; }
}

class Solution1 {

    public $count = 1;
    public $queue = [];
    /**
     * @param TreeNode $root
     * @return Integer
     */
    function maxDepth($root) {

        if ($root === null || $root->val === null) {
            return 0;
        }
        $this->queue = [$root];

        while (!empty($this->queue)) {
            $tmp = [];
            foreach ($this->queue as $treeNode) {
                $treeNode->left !== null && $tmp[] = $treeNode->left;
                $treeNode->right !== null && $tmp[] = $treeNode->right;
            }
            if (!empty($tmp)) {
                $this->count++;
            }
            $this->queue = $tmp;
        }

        return $this->count;
    }
}

class Solution2 extends Solution {

    /**
     * @param TreeNode $root
     * @return Integer
     */
    function maxDepth($root)
    {
        if ($root === null) {
            return 0;
        }

        $leftHeight = $this->maxDepth($root->left);
        $rightHeight = $this->maxDepth($root->right);

        return max($leftHeight,$rightHeight) + 1;
    }
}


$tree = new TreeNode(5);
$tree->left = new TreeNode(9);
$tree->right = new  TreeNode(11);
$tree->right->left = new  TreeNode(15);
$tree->right->right = new TreeNode(17);

$a = new Solution1();
echo $a->maxDepth($tree);

$b = new Solution2();
echo $b->maxDepth($tree);

/**
 * 解题思路:
 *   Solution1 广度优先遍历 一层一层遍历计数
 *   Solution2 递归 深度优先遍历 从叶子节点开始计数 父节点+1 遍历到跟节点
 */
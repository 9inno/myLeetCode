<?php
class Solution {

    /**
     * @param Integer $n
     * @param Integer[] $leftChild
     * @param Integer[] $rightChild
     * @return Boolean
     */
    function validateBinaryTreeNodes($n, $leftChild, $rightChild) {
        $arr = array_merge($leftChild, $rightChild);
        $tree = array_fill(0, $n -1 ,0);
        while (!empty($arr)) {
            $node = array_pop($arr);
            if ($node == -1) {
                continue;
            }
            if ($tree[$node] == 1) {
                return false;
            }else {
                $tree[$node] = 1;
                continue;
            }

        }
        $tmp = array_count_values($tree);
        if ($tmp[0] == 1) {
            return true;
        }
        return false;
    }
}
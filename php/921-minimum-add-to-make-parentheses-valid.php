<?php
class Solution {

    /**
     * @param String $S
     * @return Integer
     */
    function minAddToMakeValid($S) {

        $len = strlen($S);
        $count = 0;
        $stack = [];
        for ($i = 0 ; $i < $len; $i ++) {
            if ($S[$i] == '(') {
                array_push($stack, $S[$i]);
            } elseif ($S[$i] == ')') {
                if (empty($stack)) {
                    $count ++;
                } else {
                    array_pop($stack);
                }

            }

        }

        return $count + count($stack);
    }
}

/**
 * 解题思路：
 *   用array 模拟栈 将 左括号入栈  遍历到有括号  出栈 若是空栈 则需要补充左括号 需要补充数量 ++
 *   遍历结束 返回 需要补充的数量和 栈的元素剩余数量
 *   时间复杂度O(n)
 */
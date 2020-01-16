<?php
class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function longestValidParentheses($s) {

        $max = 0;
        $len = strlen($s);
        $stack = [-1];
        for ($i = 0; $i < $len; $i++) {


            if ($s[$i] === '(') {
                array_push($stack, $i);
            } elseif ($s[$i] === ')') {
                array_pop($stack);
                if (empty($stack)){
                    array_push($stack, $i);
                } else {
                    $max = max($max, $i- end($stack));
                }

            }


        }

        return $max;
    }
}

$a = new Solution();
echo $a->longestValidParentheses("()())");


/**
 *  解题思路：
 *   使用站存储 ( 的键  遇到） 弹出 在查看栈顶的( 的键 做减法确认距离 记录最大距离
 *   时间复杂度O(n)
 */
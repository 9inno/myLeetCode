<?php
class Solution {

    /**
     * @param Integer $n
     * @return Integer
     */
    function climbStairs($n) {
        if ($n === 1) {
            return 1;
        }
        if ($n === 2) {
            return 2;
        }

        $result = 0;
        $pre = 2;
        $prepre = 1;
        for ($i = 3; $i <= $n ; $i++) {
            $result = $pre + $prepre;
            $prepre = $pre;
            $pre = $result;
        }

        return $result;
    }
}

$a = new Solution();
$a->climbStairs(5);
/**
 * 解题思路：
 * 斐波那契数列
 * 时间复杂度 O(n)
 *
 * 其他解决方案:
 * 记忆化递归 (普通递归会栈溢出)
 * 动态规划
 * TODO
 */
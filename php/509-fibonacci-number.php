<?php
class Solution {

    /**
     * @param Integer $N
     * @return Integer
     */
    function fib($N) {
        if ($N === 1) {
            return 1;
        } elseif ($N === 0) {
            return 0;
        }

        return $this->fib($N - 1) + $this->fib($N - 2);
    }
}


$a = new Solution();
echo $a->fib(10);

/**
 *  解题思路：
 *   斐波那契数列 f(n) = f(n-1) + f(n-2)
 *   时间复杂度(2^n) 存在重复计算
 *   TODO 记忆化搜索
 */
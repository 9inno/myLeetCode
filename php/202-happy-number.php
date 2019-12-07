<?php
class Solution {

    /**
     * @param Integer $n
     * @return Boolean
     */
    function isHappy($n) {
        $len = strlen($n);
        $sum = 0;
        for ($i = 0 ; $i < $len; $i ++ ) {
            $sum += ($n % 10) ** 2;
            $n = $n / 10;
        }
        if ($sum > 4) {
            return $this->isHappy($sum);
        }

        return $sum === 1;
    }
}


/**
 *  解题思路：
 *  非快乐数 时 会进入4 16 37 58 89 145 42 20 循环
 *  所以大于4时进入递归
 *
 */
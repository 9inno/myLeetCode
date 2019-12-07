<?php
class Solution {

    /**
     * @param Integer $x
     * @return Integer
     */
    function reverse($x) {
        $reverse = 0;
        $boundary = 2 ** 31;
        $max = $boundary - 1;
        $min = -$boundary;
        while ($x != 0) {
            $tmp = $x % 10;
            $x = intval($x / 10, 0);
            $reverse = $reverse * 10 +$tmp;
        }

        if ($reverse > $max || $reverse<$min) {
            return 0;
        }

        return $reverse;
    }
}


/**
 * 解题思路：
 *  对10取模结果作为反转数的高位数字
 *  php的int型溢出会自动转为float类型
 *  所以 主动计算int的 min max
 * 时间复杂度O(log10(n))
 */
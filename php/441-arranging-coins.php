<?php
class Solution {

    /**
     * @param Integer $n
     * @return Integer
     */
    function arrangeCoins($n) {

        $i = 0;
        while (true) {
            if ($n - $i <= $i) {
                return $i;
            }
            $n -= $i;
            $i++;
        }

    }

    function arrangeCoins2($n) {
        return intval(sqrt(2) * sqrt( $n + 0.125 ) - 0.5 );
    }
}


/**
 * 解题思路：
 *   1一行一行计算
 *   2 根据公式 k(k+1)/2 =n  推算 k = sqrt(2n+ 1/4) - 1/2;  避免溢出 提出2  k = sqrt(2) * sqrt(n + 1/8) -1/2
 */
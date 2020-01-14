<?php
class Solution {

    /**
     * @param Integer $num
     * @return Integer
     */
    function addDigits($num) {

        while ($num >= 10) {

            $sum = 0;
            while (true) {
                $sum += $num % 10;
                $num = intval($num / 10);
                if ($num < 1) {
                    break;
                }

            }
            $num = $sum;
        }

        return $num;
    }

    function addDigits2($num) {
        return ($num - 1) % 9 + 1;
    }
}

$a = new Solution();
echo $a->addDigits2(10);

/**
 * 解题思路：
 *   1. 对10取模 循环相加
 *   2. 原数据-1 对9取模 +1
 */
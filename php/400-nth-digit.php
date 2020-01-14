<?php
class Solution {

    /**
     * @param Integer $n
     * @return Integer
     */
    function findNthDigit($n) {

        $i = 1;
        $sum = 0;
        while (true) {
            if ($sum +  9 * (10 ** ( $i - 1)) * $i > $n) {
                break;
            }
            $sum +=  9 * (10 ** ( $i - 1)) * $i;
            $i ++;
        }

        $div = $n - $sum;
        $p = $div % $i ;
        $target = 10 ** ($i -1) + intval($div / $i) ;
        if ($p == 0) {
            return ($target-1) % 10;
        } else {
            $target = (string)$target;
            return (int)$target[$p - 1];
        }

    }
}

$a = new Solution();
echo $a->findNthDigit(12);


/**
 * 解题思路 ：
 *   1-9 9个 10-99 90个 100-999 900个
 *   按照规律  计算位数
 * 
 */
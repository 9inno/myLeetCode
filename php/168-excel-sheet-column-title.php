<?php
class Solution {

    /**
     * @param Integer $n
     * @return String
     */
    function convertToTitle($n) {
        $n = $n - 1;
        $str = '';
        if (floor($n / 26) > 0) {
            $str .= $this->convertToTitle(floor($n / 26));
        }
        return $str . chr($n % 26 + 65);
    }
}

/**
 * 解题思路 ：
 *   26进制 ASCII码 A = 65
 */
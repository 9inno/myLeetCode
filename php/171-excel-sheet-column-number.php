<?php
class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function titleToNumber($s) {

        $tmp = 0;
        for ($i = 0, $len =strlen($s); $i < $len ; $i ++) {
            $tmp = $tmp * 26 + ord($s[$i]) - 64;
        }

        return $tmp;
    }
}


/**
 * 解题思路
 *   26进制 ASCII码 A = 65
 */
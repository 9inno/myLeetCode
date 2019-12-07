<?php
class Solution {

    /**
     * @param Integer $x
     * @return Boolean
     */
    function isPalindrome($x) {
        if ($x < 0) {
            return false;
        }
        $tmp  = $x;
        $result = 0;
        while ($tmp != 0) {
            $delivery = $tmp % 10;
            $tmp = intval($tmp / 10);
            $result = $result * 10 + $delivery;
        }

        return $result == $x;
    }
}

/**
 * 解题思路：
 * 负数肯定不是回文数，
 * 按反转整数方法 结果与元数据对比是否一致
 * 时间复杂度O(log10(n))
 *
 *
 * 优化方法： 可反转一半 与另一半对比
 * TODO
 */
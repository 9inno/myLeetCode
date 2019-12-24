<?php
class Solution {

    /**
     * @param String $str
     * @return Integer
     */
    function myAtoi($str) {

        $rules = '/^[\+\-]?\d+/';
        $flag = preg_match( $rules, trim($str), $result);
        if ($flag !== 1) {
            return 0;
        }
        return max(min((int)$result[0],(1<<31) - 1),-1<<31);
    }
}


$a = new Solution();
var_dump($a->myAtoi('    -42'));

/**
 * 解题思路：
 *   正则匹配 转为int 比较边界值
 */
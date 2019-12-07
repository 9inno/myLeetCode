<?php

class Solution {

    /**
     * @param String $haystack
     * @param String $needle
     * @return Integer
     */
    function strStr($haystack, $needle) {
        if ($needle === '') {
            return 0;
        }

        $needleLen = strlen($needle);
        $haystackLen = strlen($haystack);
        for ($i = 0; $i <= $haystackLen-$needleLen; $i++) {
            if ($haystack[$i] === $needle[0]) {
                for ($j = 1; $j < $needleLen ; $j++) {
                    if ($haystack[$i+$j] !== $needle[$j]) {
                        break;
                    }
                }
                if ($j === $needleLen) {
                    return $i;
                }
            }
        }
        return -1;
    }
}

/**
 * 解题思路：
 *
 * 遍历主字符串，从遍历到当前字符串与目标字符串第一位相同时
 * 开始遍历目标字符串 判断主字符串从当前字符串后续长度等于目标字符串长度内的字符是否与目标字符串匹配
 * 若存在不匹配的字符串则跳出 继续遍历主字符串
 * 若不存在不匹配的字符串则返回当前 遍历的key
 * 时间复杂度O(m+n)
 *
 * 优化：KMP算法
 * TODO
 */
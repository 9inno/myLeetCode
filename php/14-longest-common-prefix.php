<?php
class Solution {

    /**
     * @param String[] $strs
     * @return String
     */
    function longestCommonPrefix($strs) {
        $count = count($strs);
        $result = '';
        $j = 0;
        while ($count) {
            $tmp = '';
            for ($i=0 ; $i <$count; $i ++) {
                if ($tmp === '') {
                    $tmp = $strs[$i][$j];
                }
                if (!isset($strs[$i][$j]) || $tmp != $strs[$i][$j]  ) {
                    return $result;
                }
            }

            $result = $result.$tmp;
            $j++;
        }
        return  $result;
    }
}

/**
 * 解题思路：
 * 遍历数组元素的同一位置的字符串
 * 当遍历完最短元素 或 同一位置的字符串不相等时结束
 * 时间复杂度 O(mn) m为数组元素最短字符串的长度 n为数组元素个数
 */
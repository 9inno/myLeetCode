<?php
class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function lengthOfLongestSubstring($s) {
        $len = strlen($s);
        $hash = [];
        $count = 0;
        for ($i = 0; $i < $len; $i++) {
            if (in_array($s[$i], $hash)) {
                while (true) {
                    $unshift = array_shift($hash);
                    if ($unshift === $s[$i]) {
                        break;
                    }
                }
            }
            $hash[] = $s[$i];
            $count = max(count($hash), $count);
        }

        return $count;
    }
}

/**
 * 解题思路：
 * 遍历字符串 不重复的丢到数组里
 * 遇到重复的字符串，
 * 并从数组开头弹出元素，直到弹出当前重复元素为止
 *
 * 计算当前数组元素个数， 与当前计数 取较大值
 * 时间复杂度O(n)
 */
<?php
class Solution {

    /**
     * @param String $s
     * @param String $t
     * @return Boolean
     */
    function isAnagram($s, $t) {
        $length = strlen($s);
        if ($length !== strlen($t)) {
            return false;
        }

        $map = [];

        for ($i = 0; $i < $length; $i++) {
            $map[$s[$i]]++;
            $map[$t[$i]]= isset($map[$t[$i]]) ? $map[$t[$i]] - 1 : -1;
        }

        foreach ($map as $value) {
            if ($value !== 0) {
                return false;
            }
        }

        return true;
    }
}

$a = new Solution();
var_dump($a->isAnagram('anagram', 'nagaram'));

/**
 *  解题思路：
 *   同时遍历 两个字符串 放入哈希map 一个+1 一个-1
 *   遍历结束后 遍历哈希map 如果有不为零的值 则无效
 *   时间复杂度O(n)
 */
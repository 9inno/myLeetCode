<?php
class Solution {

    /**
     * @param String[] $strs
     * @return String[][]
     */
    function groupAnagrams($strs) {
        $result = [];
        foreach ($strs as $str) {
            $tmp = str_split($str);
            sort($tmp);
            $result[implode('', $tmp)][] = $str;
        }
        return array_values($result);
    }
}

$a = new Solution();
print_r($a->groupAnagrams(["eat","tea","tan","ate","nat","bat"]));

/**
 * 解题思路：
 *    将字符串 转成数组 排序后做hashmap
 *    
 */
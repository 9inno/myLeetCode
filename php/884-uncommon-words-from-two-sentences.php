<?php
class Solution {

    /**
     * @param String $A
     * @param String $B
     * @return String[]
     */
    function uncommonFromSentences($A, $B) {
        $arrA = explode(' ', $A);
        $arrB = explode(' ', $B);
        $tmp = array_count_values(array_merge($arrA, $arrB));
        $result = [];
        foreach ($tmp as $key => $value) {
            if ($value == 1) {
                $result[] = $key;
            }
        }
        return $result;
    }
}

$a = new Solution();
print_r($a->uncommonFromSentences("apple apple","banana"));

/**
 * 解题思路
 *  统计所有单词的 出现的次数
 */
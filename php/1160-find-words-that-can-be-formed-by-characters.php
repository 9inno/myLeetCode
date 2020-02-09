<?php
class Solution {

    /**
     * @param String[] $words
     * @param String $chars
     * @return Integer
     */
    function countCharacters($words, $chars) {
        $arr = array_count_values(str_split($chars));
        $result = 0;
        foreach ($words as $word) {
            $flag = true;
            foreach (array_count_values(str_split($word)) as $key => $value) {
                if (!isset($arr[$key]) || $arr[$key] < $value) {
                    $flag = false;
                }
            }
            if ($flag) {
                $result += strlen($word);
            }
        }

        return $result;
    }
}
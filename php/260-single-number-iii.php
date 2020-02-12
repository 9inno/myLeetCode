<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer[]
     */
    function singleNumber($nums) {
        $arr = array_count_values($nums);
        $result = [];
        foreach ($arr as $key => $value) {
            if ($value == 1) {
                $result[] = $key;
            }
        }
        return $result;
    }
}
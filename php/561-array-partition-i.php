<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function arrayPairSum($nums) {
        sort($nums);
        $result = 0;
        $count = count($nums);
        for ($i = 0; $i < $count; $i += 2) {
            $result += $nums[$i];
        }

        return $result;
    }
}


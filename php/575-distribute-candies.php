<?php
class Solution {

    /**
     * @param Integer[] $candies
     * @return Integer
     */
    function distributeCandies($candies) {
        $arr = array_count_values($candies);
        if (count($arr) > (count($candies) / 2)) {
            return (count($candies) / 2);
        }
        return count($arr);
    }
}
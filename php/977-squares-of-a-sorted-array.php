<?php
class Solution {

    /**
     * @param Integer[] $A
     * @return Integer[]
     */
    function sortedSquares($A) {
        $result = [];
        foreach ($A as $value) {
            $result[] = $value ** 2;
        }
        sort($result);
        return $result;
    }
}
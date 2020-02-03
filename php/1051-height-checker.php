<?php
class Solution {

    /**
     * @param Integer[] $heights
     * @return Integer
     */
    function heightChecker($heights) {
        $tmp = $heights;
        sort($heights);
        $count = count($heights);
        $result = 0;
        for ($i = 0; $i < $count; $i++) {
            if ($tmp[$i] != $heights[$i]) {
                $result++;
            }
        }

        return $result;
    }
}
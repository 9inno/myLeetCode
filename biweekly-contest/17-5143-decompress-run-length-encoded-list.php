<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer[]
     */
    function decompressRLElist($nums) {

        $count = count($nums) / 2;

        $result = [];
        for ($i = 0 ; $i < $count; $i++ ) {
            list($a, $b) = [$nums[$i * 2], $nums[$i * 2 + 1]];

            for ($j = 0; $j < $a; $j ++) {
                $result[] = $b;
            }
        }

        return $result;

    }
}
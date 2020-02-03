<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer
     */
    function search($nums, $target) {

        $left = 0;
        $right = count($nums);
        while ($left <= $right) {
            $mid = intval(($left + $right) / 2);
            if ($nums[$mid] ==$target) {
                return $mid;
            }
            if ($nums[$mid] > $target) {
                $right = $mid - 1;
            }
            if ($nums[$mid] < $target) {
                $left = $mid + 1;
            }

        }

        return -1;
    }
}

$a = new Solution();
echo $a->search([-1,0,3,5,9,12], 9);
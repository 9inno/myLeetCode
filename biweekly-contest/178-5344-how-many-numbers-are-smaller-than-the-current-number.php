<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer[]
     */
    function smallerNumbersThanCurrent($nums) {
        return array_map(function ($num) use($nums) {
            $count = 0;
            foreach ($nums as $i) {
                if ($num> $i) {
                    $count++;
                }
            }
            return $count;
        },$nums);
    }
}

$a = new Solution();
print_r($a->smallerNumbersThanCurrent([8,1,2,2,3]));
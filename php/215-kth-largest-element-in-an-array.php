<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $k
     * @return Integer
     */
    function findKthLargest($nums, $k) {
        $arr = array_unique($nums);
        rsort($arr);
        return $arr[$k - 1];
    }
}

$a = new Solution();
$a->findKthLargest([3,2,3,1,2,4,5,5,6],4);
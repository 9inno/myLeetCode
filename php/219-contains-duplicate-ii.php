<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $k
     * @return Boolean
     */
    function containsNearbyDuplicate($nums, $k) {
        $map = [];
        $count = count($nums);
        for($i = 0; $i < $count; $i++) {
            if (isset($map[$nums[$i]])) {
                return true;
            }
            $map[$nums[$i]] = true;
            if (count($map) > $k) {
                unset($map[$nums[$i - $k]]);
            }
        }
        return false;
    }
}

$a = new Solution();
var_dump($a->containsNearbyDuplicate([1,2,3,1,2,3], 2));

/**
 * 解题思路:
 *    利用数组做hashmap 维护一个k大小的 hashmap 判断是否存在
 *    时间复杂度O(n)
 */
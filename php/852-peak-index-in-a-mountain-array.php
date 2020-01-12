<?php
class Solution {

    /**
     * @param Integer[] $A
     * @return Integer
     */
    function peakIndexInMountainArray($A) {
        $count = count($A);
        for ($i = 1 ; $i < $count; $i ++) {
            if ($A[$i-1] < $A[$i] && $A[$i+1] < $A[$i] ) {
                return $i;
            }
        }

        return 0;
    }

    function peakIndexInMountainArray2($A) {
        $right = count($A) - 1;
        $left = 0;
        while ($left <= $right) {
            $mid = intval(($right + $left) / 2);
            if ($A[$mid - 1] < $A[$mid] && $A[$mid + 1] < $A[$mid]) {
                return $mid;
            }
            if ($A[$mid - 1] < $A[$mid] && $A[$mid] < $A[$mid +1]) {
                $left = $mid;
            }
            if ($A[$mid - 1] > $A[$mid] && $A[$mid] > $A[$mid +1]) {
                $right = $mid;
            }
        }

        return 0;
    }
}


$a = new Solution();
echo $a->peakIndexInMountainArray2([0,1, 0]);

/**
 * 解题思路：
 *   遍历数组
 *   时间复杂度O(n)
 *   二分查找法
 *   时间复杂度O(logn)
 */
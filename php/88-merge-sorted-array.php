<?php
class Solution {

    /**
     * @param Integer[] $nums1
     * @param Integer $m
     * @param Integer[] $nums2
     * @param Integer $n
     * @return NULL
     */
    function merge(&$nums1, $m, $nums2, $n) {

        $i = $m - 1;
        $j = $n - 1;
        $p = $m + $n -1;
        while($i >= 0 && $j >= 0) {
            $nums1[$p --] = $nums1[$i] > $nums2[$j] ? $nums1[$i --] : $nums2[$j --];
        }

        while ($i < 0 && $j >= 0) {
            $nums1[$p --] = $nums2[$j --];
        }

    }
}


$nums1 =[1,2,3,0,0,0];
$m = 3;
$nums2 =[2,5,6];
$n = 3;
$a = new Solution();
$a->merge($nums1, $m, $nums2, $n);
print_r($nums1);

/**
 * 解题思路：
 *   双指针 从数组最后 向前遍历
 *   p指针用于 添加元素的位置
 *   时间复杂度O(m+n)
 */
<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function removeDuplicates(&$nums) {
        $count = count($nums);
        $j = 0;
        for ($i = 0; $i < $count; $i ++ ) {
            if ($nums[$j] !== $nums[$i]) {
                $nums[$j + 1] = $nums[$i];
                $j ++;
            }
        }
        return $j + 1;
    }
}

$arr = [0,0,1,1,1,2,2,3,3,4];
$a = new Solution();
$a->removeDuplicates($arr);
print_r($arr);

/**
 *  解题思路：
 *   因数组是有序数组  那么重复的元素是相邻的
 *   快慢指针遍历数组 慢指针指向开始元素 快指针向后遍历
 *   遇到快指针元素 与慢指针元素不等时 将快指针元素 赋值给慢指针下一元素 并将慢指针指向下一元素
 *   时间复杂度O(n)
 */
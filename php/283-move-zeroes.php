<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @return NULL
     */
    function moveZeroes(&$nums) {
        $index = 0;
        $count = count($nums);
        for ($i = 0; $i < $count; $i++) {
            if ($nums[$i] != 0) {
                $nums[$index] = $nums[$i];
                if ($index != $i) {
                    $nums[$i] = 0;
                }
                $index++;
            }
        }
    }
}

$b = [1,1,0,3,12];
$a = new Solution();
$a->moveZeroes($b);
print_r($b);

/**
 * 解题思路 ：
 *   双指针 遍历数组 快慢指针同时出发 遇到0时 快指针继续遍历 慢指针停留在0的位置
 *   然后将快指针的值 赋给慢指针 快指针赋值0
 *   时间复杂度O(n)
 */
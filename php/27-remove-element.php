<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $val
     * @return Integer
     */
    function removeElement(&$nums, $val) {

        $count = count($nums);
        $j = 0;
        for ($i = 0; $i < $count; $i++) {
            if ($nums[$i] != $val) {
                $nums[$j] = $nums[$i];
                $j++;
            }
        }
        return $j;
    }
}

$arr = [0,1,2,2,3,0,4,2];
$a = new Solution();
echo $a->removeElement($arr, 2);
print_r($arr);

/**
 *  解题思路
 *   快慢指针 快慢指针一起出发 慢指针遇到 目标值停下 快指针继续向后
 *   遇到不是目标值的元素 赋值与慢指针 慢指针指向下一元素
 *   时间复杂度O(n)
 */
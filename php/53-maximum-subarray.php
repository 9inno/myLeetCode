<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function maxSubArray($nums) {
        $count = count($nums);
        $maxNum = $nums[0];
        for($i = 1 ; $i < $count; $i++) {
                $nums[$i] = max($nums[$i], $nums[$i] + $nums[$i - 1]);
                $maxNum = max($maxNum, $nums[$i]);
        }
        return $maxNum;
    }


}

//$arr = [-2,1];
$arr = [-2,1,-3,4,-1,2,1,-5,4];
$a = new Solution();
echo $a->maxSubArray($arr);

/**
 *  解题思路：
 *   遍历数组 当前元素 可能为 1. 连续子数组的开头 2.连续子数组中间活结尾
 *   单讨论$nums[$i] = $num[$i] + $nums[$i-1]  即 当前元素赋值为前序元素的和
 *   当当前元素大于 与前序元素和时 说明 前序来连续子数组的和不是最大
 *   则以当前元素为 连续子数组的开头 继续向后计算
 *   时间复杂度 O(n)
 */
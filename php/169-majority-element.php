<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function majorityElement($nums) {
        $map = [];
        $target = intval(count($nums) / 2 );
        foreach ($nums as $num) {
            if (isset($map[$num])) {
                $map[$num]++;
                if ($map[$num] > $target) {
                    return $num;
                }
            } else {
                $map[$num] = 1;
            }
        }
        return $nums[0] ?? 0;
    }
}

$a = new Solution();
echo $a->majorityElement([1,1,2,2,3,3,3]);



/**
 * 解题思路 ：
 *   遍历数组 记录出现次数 返回出现次数大于 n/2的数字
 *   时间复杂度O(n)
 */
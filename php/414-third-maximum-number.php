<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function thirdMax($nums) {
        $first = null;
        $second = null;
        $third = null;

        foreach ($nums as $num) {
            if ($num === $first || $num === $second || $num === $third) {
                continue;
            }
            if ($num > $first) {
                $third = $second;
                $second = $first;
                $first = $num;
            }
            if ($num < $first && $num > $second) {
                $third = $second;
                $second = $num;
            }
            if ($num < $second && $num >= $third) {
                $third = $num;
            }
        }

        return $third !== null? $third : $first;
    }
}

//test
$a = new Solution();
echo $a->thirdMax([3,3,4,3,4,3,0,3,3]);


/**
 * 解题思路 ：
 *    遍历数组 当前元素与 第一大第二大第三大元素 对比
 *   若满足条件则 使当前元素成为 第几大 依次向下替换
 *   时间复杂度O(n)
 */
<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    function twoSum($nums, $target) {
        $tmp = [];
        foreach ($nums as $key => $value) {
            if (isset($tmp[($target-$value)])) {
                return [$tmp[($target-$value)], $key];
            }else {
                $tmp[$value] = $key;
            }
        }
        return [];
    }
}


/**
 * 解题思路：
 * 遍历数组 与 目标值做差，  确认差值是否在新数组中存在简直相同键
 * 若在 则返回当前key与新数组差值为键的值，若不在则放入新数组 值为新数组的键 键为新数组的值
 * 时间复杂度O(n)
 */

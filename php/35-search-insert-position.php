<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer
     */
    function searchInsert($nums, $target) {
        $result = null;
        foreach ($nums as $key => $num) {
            if ($num >= $target) {
                $result = $key;
                break;
            }
        }
        if ($result === null) {
            $result = count($nums);
        }
        return $result;
    }
}


/**
 *
 * 解题思路：
 *  遍历数组 若当前元素大于等于目标值时
 *  返回当前键值
 *  若遍历结束 不存在结果 则插入数组最后
 *  时间复杂度O(n)
 *
 * 优化：
 *   二分查找法
 *   TODO
 */
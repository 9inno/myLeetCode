<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @return Boolean
     */
    function containsDuplicate($nums) {

        while (!empty($nums)) {
            $element = array_pop($nums);
            if (in_array($element,$nums)) {
                return true;
            }
        }

        return false;
    }


    function containsDuplicate2($nums) {

        $map = [];
        foreach ($nums as $num) {
            if (isset($map[$num])) {
                return true;
            }else {
                $map[$num] = true;
            }
        }
        return false;
    }
}

/**
 * 解题思路
 *   1 弹出数组元素 查看数组是否存在相同元素
 *   2哈希查找
 *   时间复杂度O(n)
 */
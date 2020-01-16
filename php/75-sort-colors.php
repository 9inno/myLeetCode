<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @return NULL
     */
    function sortColors(&$nums) {
        $map = [
            0 => 0,
            1 => 0,
            2 => 0,
        ];
        foreach ($nums as $num) {
            $map[$num] ++;
        }

        $i = 0;
        foreach ($map as $key => $value) {
            for ($j = 0 ; $j < $value; $j++) {
                $nums[$i] = $key;
                $i ++;
            }
        }
    }
}

/**
 * 解题思路：
 *   计数排序
 *   遍历数组 记录 每项个数 再遍历个数赋值到原数组
 */
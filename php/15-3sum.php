<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer[][]
     */
    function threeSum($nums) {

        $count = count($nums);
        sort($nums);
        $result = [];
        $done = [];
        for ($i = 0; $i < $count; $i++) {
            if (isset($done[$nums[$i]])) {
                continue;
            }
            $map = [];
            for ($j = $i + 1; $j < $count; $j++) {
                if (!isset($map[-($nums[$i] + $nums[$j])])) {
                    $map[$nums[$j]] = true;
                } else {
                    $tmp = -($nums[$i] + $nums[$j]);
                    $result["$nums[$i],$tmp,$nums[$j]"] = [$nums[$i], $tmp, $nums[$j]];
                }
            }
            $done[$nums[$i]] = true;
        }
        return array_values($result);
    }
}

$a = new Solution();
print_r($a->threeSum($arr = [0,0,0,0]));

/**
 * 解题思路 ：
 *   寻找三个数加和等于零  固定一个数字 变成 寻找两个数的加和的相反数等于固定的数字
 *   因答案中不可以包含重复的三元组 去重比较麻烦。  先将数组排序 重复的数字将挨在一起 已经固定过的数字 直接跳过
 *   结果集 有特殊情况 000 000 的情况  所以用数组的键值 来去重
 *   时间复杂度 O(n^2)
 */


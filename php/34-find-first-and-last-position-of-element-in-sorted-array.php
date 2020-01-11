<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    function searchRange($nums, $target) {
        $result = [-1, -1];
        if (empty($nums)) {
            return $result;
        }
        $left = 0;
        $right = count($nums) - 1;
        $p = null;

        while ($left <= $right) {
            $mid = intval(($left + $right) / 2);
            if ($nums[$mid] === $target) {
                $p = $mid;
                break;
            } elseif($nums[$mid] > $target) {
                $right = $mid - 1;
            } elseif ($nums[$mid] < $target) {
                $left = $mid + 1;
            }
        }

        if ($p !== null) {
            list($result[0], $result[1]) = [$p, $p];
            while (true) {
                if ($nums[$result[0] - 1] === $target) {
                    $result[0] --;
                }
                if ($nums[$result[1] + 1] === $target) {
                    $result[1] ++;
                }
                if ($nums[$result[1] + 1] !== $target && $nums[$result[0] - 1] !== $target) {
                    break;
                }
            }
        }

        return $result;
    }
}


$a = new Solution();
print_r($a->searchRange([5,7,7,8,8,10], 8));


/**
 * 解题思路：
 *   二分法查找target 后向前后扫描 记录target相同的开始和结束位置
 */
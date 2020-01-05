<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function rob($nums) {
        $count = count($nums);
        if ($count == 0){
            return 0;
        }elseif ($count == 1) {
            return $nums[0];
        }elseif ($count == 2) {
            return max($nums[0], $nums[1]);
        }
        $first[0] = $nums[0];
        $first[1] = $nums[0];
        $noFirst[0] = 0;
        $noFirst[1] = $nums[1];
        for ($i = 2; $i < $count; $i++) {
            $first[$i] = max($first[$i-2] + $nums[$i], $first[$i-1]);
            $noFirst[$i] = max($noFirst[$i-2] + $nums[$i], $noFirst[$i-1]);
        }
        return max($first[$count-2], $noFirst[$count-1]);
    }
}


$a = new Solution();
print_r( $a->rob([1,2,3,1]));


/**
 * 解题思路 ：
 *   动态规划 OPT($i) = 选i节点 OPT($i-2) + $num[$i]
 *                      不选i节点 OPT($i -1)
 *   由于 首尾不能同时选 所以分解成 第一个位置选或不选两种情况进行dp
 *   由于选了第一个位置就不能选最后一个位置 所以 选第一个位置的结果应该为count-2
 *   时间复杂度O(n)
 */
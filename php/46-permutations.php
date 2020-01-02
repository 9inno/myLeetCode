<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer[][]
     */
    function permute($nums) {
        $count = count($nums);
        $result = [];

        $this->recursion($result, $nums, 0, $count, []);
        return $result;
    }

    public function recursion(&$result, $nums, $n, $count, $tmp) {
        if ($count === 0 ) {
            return;
        }
        if ($count === $n) {
            array_push($result, $tmp);
            return;
        }
        for ($i = $n; $i < $count; $i ++) {
            $p = $nums[$i];
            $nums[$i] = $nums[$n];
            $nums[$n] = $p;
            $this->recursion($result, $nums, $n + 1, $count, array_merge($tmp, [$nums[$n]]));
            $p = $nums[$i];
            $nums[$i] = $nums[$n];
            $nums[$n] = $p;
        }
    }
}

$a = new Solution();
print_r($a->permute([1,2,3]));


/**
 * 解题思路：
 *  遍历数组 当前元素与 指针元素互换位置
 *  满足条件计入结果
 *  互换位置回溯回之前的数组
 *  指针向后移动
 *  时间复杂度O(n!)
 */
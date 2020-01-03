<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer[][]
     */
    function subsets($nums) {
    $result = [];
    $this->recursion($result, $nums, 0, []);
    return $result;
}


    function recursion(&$result, $nums, $n, $tmp) {
        if ($n >= count($nums)) {
            $result[] = $tmp;
            return;
        }

        $this->recursion($result, $nums, $n+1, array_merge($tmp, [$nums[$n]]));
        $this->recursion($result, $nums, $n+1, $tmp);
    }
}


$a = new Solution();
print_r($a->subsets([1,2,3]));


/**
 * 解题思路
 *   递归遍历数组 每个元素 可以有 选和不选的情况
 *   遍历到数组结束 返回结果
 *   时间复杂度O(n)
 */
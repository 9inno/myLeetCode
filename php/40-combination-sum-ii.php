<?php
class Solution {

    /**
     * @param Integer[] $candidates
     * @param Integer $target
     * @return Integer[][]
     */
    function combinationSum2($candidates, $target) {
        $result = [];
        sort($candidates);
        $this->recursion($result, $candidates, $target, 0, []);
        return $result;
    }

    function recursion(&$result, $candidates, $target, $n, $tmp) {

        if ($target < 0 || $n > count($candidates)) {
            return;
        }
        if ($target === 0) {
            $result[] = $tmp;
        }

        $map = [];

        for ($i = $n ; $i < count($candidates); $i ++) {
            if (in_array($candidates[$i], $map)) {
                continue;
            }
            $this->recursion($result, $candidates, $target - $candidates[$i], $i + 1, array_merge($tmp, [$candidates[$i]]));
            $map[] = $candidates[$i];
        }

    }
}


$a = new Solution();
print_r($a->combinationSum2([10,1,2,7,6,1,5],8));


/**
 *  解题思路：
 *   排序数组后  递归遍历数组  递归一层 指针向后移动一位
 *   当前递归内 存在已经 选过的相同的数则跳过
 *
 */
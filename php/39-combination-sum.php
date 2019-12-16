<?php
class Solution {

    /**
     * @param Integer[] $candidates
     * @param Integer $target
     * @return Integer[][]
     */
    function combinationSum($candidates, $target) {
        if ($target < 0) {
            return [];
        }
        if ($target === 0) {
            return  [[]];
        }
        $result = [];
        sort($candidates);
        foreach ($candidates as $candidate) {
            $difference = $target - $candidate;

            if ($difference === 0) {
                return array_merge($result,[[$candidate]]);
            } elseif ( $difference !== 0) {
                $tmp = $this->combinationSum($candidates, $difference);
                if (!empty($tmp)) {
                    foreach ($tmp as &$item) {
                        array_unshift($item, $candidate);
                    }
                    unset($item);
                    $result = array_merge($result, $tmp);
                }

            }
            array_shift($candidates);
        }
        return $result;
    }
}


$a = new Solution();
print_r($a->combinationSum([4,2,8], 8));

/**
 * 解题思路：
 *   深度优先搜索  排序剪枝
 *   因时排序过后的数组 递归深度增加时 递归传入的数组元素逐渐减少  更深层的值不会比上层值小 否则会出现位置不同的重复结果
 *
 */
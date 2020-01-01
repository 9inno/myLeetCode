<?php

class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer
     */
    function threeSumClosest($nums, $target) {
        $count = count($nums);
        sort($nums);
        $min = null;
        $result = null;
        for ($i = 0; $i < $count - 2 ; $i ++) {
            $j = $i + 1;
            $k = $count - 1;
            while ($k > $j) {
                $tmp = $target - ($nums[$i] + $nums[$j] + $nums[$k]);
                $abs = abs($tmp);
                if ($min === null) {
                    $min = $abs;
                    $result = $nums[$i] + $nums[$j] + $nums[$k];
                }
                if ($abs < $min) {
                    $min = $abs;
                    $result = $nums[$i] + $nums[$j] + $nums[$k];
                }

                if ($tmp > 0) {
                    $j ++;
                } else {
                    $k --;
                }
            }
        }
        return $result;
    }
}


$a = new Solution();
echo $a->threeSumClosest($arr = [-1,2,1,-4], 1);


/**
 * 解题思路 ：
 *   先对数组排序  然后固定一个数字  然后从剩下的数字 两边开始取值相加
 *   与target 做差取绝对值比较大小
 *   时间复杂度  O(n^2)
 */
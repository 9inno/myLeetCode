<?php
class Solution {

    /**
     * @param Integer[] $prices
     * @return Integer
     */


    function maxProfit($prices) {
        $count = count($prices);
        $max = 0;
        $buy = $prices[0];
        for ($i = 0; $i < $count; $i ++ ) {
            if ($buy > $prices[$i]) {
                $buy = $prices[$i];
            } else {
                $max = max($max, $prices[$i] -$buy);
            }
        }
        return $max;
    }


    function maxProfit1($prices) {
        $count = count($prices);
        $max = 0;
        for ($i = 0; $i < $count; $i ++) {
            for ($j = $i + 1; $j < $count; $j ++) {
                $max = max($max,$prices[$j] - $prices[$i]);
            }
        }

        return $max;
    }
}


$arr = [7,17,5,3,6,4];
$a = new Solution();
echo  $a->maxProfit($arr);

/**
 * 解题思路 ：
 *   遍历数组 从第一天买入 尝试每天卖出 记录最大收益
 *   遇到元素比购买价格小的时候 替换购买 继续向后遍历 尝试每天卖出 记录最大收益
 *   遍历结束 最大收益即为结果
 *   时间复杂度O(n)
 */
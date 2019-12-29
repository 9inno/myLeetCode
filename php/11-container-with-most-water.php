<?php
class Solution {

    /**
     * @param Integer[] $height
     * @return Integer
     */
    function maxArea($height) {
        $count = count($height);
        $i = 0;
        $j = $count -1;
        $max = 0;

        while (true) {

            $max =max($max,min($height[$i], $height[$j]) * ($j - $i));

            if ($i == $j -1) {
                break;
            }

            if ($height[$j] < $height[$i]) {
                $j --;
            } else {
                $i ++;
            }


        }

        return $max;
    }
}
$arr = [1,8,6,2,5,4,8,3,7];
$a = new Solution();
echo $a->maxArea($arr);

/**
 * 解题思路：
 *   双指针 两个数形成的面积 为小的数字与距离的乘积
 *   所以挪动较小数的指针 继续计算面积  记录面积的最大值
 *   时间复杂度O(n)
 */
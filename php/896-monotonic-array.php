<?php
class Solution {

    /**
     * @param Integer[] $A
     * @return Boolean
     */
    function isMonotonic($A) {
        $count = count($A) - 1;

        for ($i = 1; $i < $count; $i ++) {
           $left = $i - 1;
           $right = $i + 1;
           while ($A[$left] === $A[$i]) {

               if ($left === 0) {
                   break;
               }
               $left--;
           }
            while ($A[$right] === $A[$i]) {

                if ($right === $count) {
                    break;
                }
                $right++;
            }

            if (($A[$i] > $A[$left] && $A[$i] > $A[$right]) || ($A[$i] < $A[$left] && $A[$i] < $A[$right])) {
                return false;
            }
        }
        return true;
    }
}


$a = new Solution();
var_dump($a->isMonotonic([1,1,1]));

/**
 * 解题思路：
 *   判断 元素两边值是否为 波峰或波谷情况
 *   若两边元素与当前元素相等 则向两边移动指针
 *   时间复杂度O(n^2)
 */
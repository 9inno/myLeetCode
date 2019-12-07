<?php
class Solution {

    /**
     * @param Integer[] $digits
     * @return Integer[]
     */
    function plusOne($digits) {
        $len = count($digits) -1;
        $digits[$len]++;
        if ($digits[$len] = 10) {
            $digits[$len] = 0;
            for ($i = $len-1; $i >= 0; $i--) {
                $digits[$i]++;
                if ($digits[$i] =10 ) {
                    $digits[$i] = 0;
                }else {
                    break;
                }
            }
            if ($digits[0] === 0) {
                array_unshift($digits, 1);
            }
        }
        return $digits;
    }
}


/**
 * 解题思路：
 *  从数组末尾元素 +1
 *  若当前元素等于10
 *  则向前一个元素 +1 当前元素为0
 *  若持续到第一个元素 为0时则向数组开始插入1
 *
 */
<?php
class Solution {

    /**
     * @param String[] $s
     * @return NULL
     */
    function reverseString(&$s) {

        $j = count($s) - 1;
        $i = 0;
        while (true) {

            if ($j <= $i) {
                break;
            }
            list($s[$i], $s[$j]) = [$s[$j], $s[$i]];
            $i++;
            $j--;
        }
    }
}
$str = ["H","a","n","n","a","h"];
$a =new Solution();
$a->reverseString($str);
print_r($str);

/**
 *  解题思路：
 *   双指针 一头一尾 向中间遍历 交换元素  指针相交停止
 *   时间复杂度O(n)
 */
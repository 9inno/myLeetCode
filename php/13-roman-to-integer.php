<?php
class Solution {

    private $map = [
        'I' => 1,
        'V' => 5,
        'X' => 10,
        'L' => 50,
        'C' => 100,
        'D' => 500,
        'M' => 1000
    ];
    /**
     * @param String $s
     * @return Integer
     */
    function romanToInt($s) {
        $len =strlen($s);
        $result = 0;
        for ($i = 0; $i < $len; $i++) {
            $current = $this->map[$s[$i]];
            $next = $this->map[$s[$i+1]];

            if ($current< $next) {
                $result = $result - $current;
            } else {
                $result = $result + $current;
            }
        }

        return $result;
    }
}


$a = new Solution();
echo  $a->romanToInt('III');

/**
 * 解题思路：
 *   遍历字符串 若当前字符串的值小于下一个字符串的值则减去当前的值
 *   若大于则加上当前的值
 *   时间复杂度O(n)
 */
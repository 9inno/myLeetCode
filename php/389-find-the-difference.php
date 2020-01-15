<?php
class Solution {

    /**
     * @param String $s
     * @param String $t
     * @return String
     */
    function findTheDifference($s, $t) {
        $len = strlen($s);
        $map = [];
        for ($i = 0; $i < $len ; $i++) {
            $map[$s[$i]] ++;
        }

        for ($j = 0; $j < $len + 1; $j ++) {
            if (!isset($map[$t[$j]])) {
                return $t[$j];
            }else {
                $map[$t[$j]] --;
                if ($map[$t[$j]] <0) {
                    return $t[$j];
                }
            }
        }
    }
}

/**
 * 解题思路：
 *   将原始字符串做成map 并计数
 *   遍历重排字符串 对应map -1 若不在map中或map中元素小于0 则是新添加的字符串
 *   时间复杂度O(n)
 */
<?php
class Solution {

    /**
     * @param String $moves
     * @return Boolean
     */
    function judgeCircle($moves) {
        $x = 0;
        $y = 0;
        for ($i = 0; $i < strlen($moves); $i ++) {
            if ($moves[$i] === 'U') {
                $y++;
            } elseif ($moves[$i] === 'D') {
                $y--;
            } elseif ($moves[$i] === 'L') {
                $x--;
            } elseif ($moves[$i] === 'R') {
                $x++;
            }
        }
        return $x === 0 && $y ===0;
    }
}


$a = new Solution();
var_dump($a->judgeCircle("LL"));

/**
 *  解题思路：
 *   遍历字符串 移动目标 记录 坐标x y 位置
 *   时间复杂度O(n)
 */
<?php
class Solution {

    /**
     * @param String $s
     * @return Boolean
     */
    function isValid($s) {
        $stack = [];
        $match = [
            '[' => ']',
            '(' => ')',
            '{' => '}'
        ];
        for ($i=0; $i<strlen($s); $i++) {
            if ($s[$i] === '(' || $s[$i] === '{' || $s[$i] === '['  ) {
                $stack[] = $s[$i];
            } else {
                if ($match[array_pop($stack)] !== $s[$i]) {
                    return false;
                }
            }
        }
        return count($stack) == 0;
    }
}

/**
 * 解题思路：
 * 用php数组当作栈
 * 遍历字符串，左括号压入栈，
 * 碰到右括号从栈弹出元素 判断是否匹配
 * 遍历结束 查看栈中是否还有元素
 * 若无则有效 若有则无效
 * 时间复杂度O(n)
 */
<?php
class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function lengthOfLastWord($s) {
        $len = strlen($s);
        $i = $len - 1;
        while($s[$i] === ' ') {
            $i--;
            $len--;
        }
        while ($i >= 0) {
            if ($s[$i] === ' ') {
                break;
            }
            $i--;
        }
        return $len - $i - 1;
    }
}

$a = new Solution();
echo $a->lengthOfLastWord("HelloWorld");

/**
 *  解题思路：
 *    从后往前遍历字符串 记录单词长度
 *    注意 以空格结尾的字符串
 */
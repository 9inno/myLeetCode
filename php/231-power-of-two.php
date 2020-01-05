<?php
class Solution {

    /**
     * @param Integer $n
     * @return Boolean
     */
    function isPowerOfTwo($n) {
        return $n > 0 && ($n & ($n - 1)) == 0;
    }
}


$a = new Solution();
var_dump($a->isPowerOfTwo(8));


/**
 * 解题思路 ：
 *   位运算 2的幂 二进制 最高位为1 其他位为0
 *   2的幂 -1 最高位为0 其他位为1 & 运算恒等于0
 *
 */
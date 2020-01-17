<?php
class Solution {

    /**
     * @param Integer $N
     * @param Integer $K
     * @return Integer
     */
    function kthGrammar($N, $K) {
        $binary =(string) decbin($K - 1);
        return  substr_count($binary,'1')% 2;
    }
    function kthGrammar2($N, $K) {
        if($N == 1) {
            return 0;
        }
        if($K <= 1 << $N - 2) {
            return $this->kthGrammar2($N - 1, $K);
        }
        return $this->kthGrammar2($N - 1, $K - (1<< $N -2)) ^ 1;
    }
}


$a = new Solution();
echo $a->kthGrammar2(10, 5);


/**
 *   解题思路：
 *
 *      0
 *      0|1
 *      01|10
 *      0110|1001
 *    递归： 按照规律 每行分为两部分 前部分时候半部分的反转  1 -> 0  0->1
 *      如果K在第二部分 找K-(1 << N -2) 在第一部分的答案 然后反转
 *
 *    规律 ：把下标 K 写成二进制的形式，如果 K 在第二部分，那么二进制形式的第一个比特位一定是 1。据此可以将方法三继续简化，实际上翻转的次数等于 K-1 二进制表示形式中 1 出现的个数。
 *
 */
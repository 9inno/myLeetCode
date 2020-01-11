<?php
class Solution {

    /**
     * @param Integer $x
     * @return Integer
     */
    function mySqrt($x) {
       if ($x == 0 ) return 0;
       $last = 0;
       $result = 1;
       while ($last != $result) {
           $last = $result;
           $result = ($result+$x/$result)/2;
       }
       return intval($result);
    }
}


/**
 * 解题思路：
 *   牛顿迭代法
 */